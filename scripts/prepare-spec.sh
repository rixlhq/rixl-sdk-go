#!/usr/bin/env bash

set -euo pipefail

SCRIPT_DIR="$(cd -- "$(dirname -- "${BASH_SOURCE[0]}")" && pwd)"
ROOT_DIR="$(cd -- "${SCRIPT_DIR}/.." && pwd)"
DEFAULT_INPUT_SPEC="${ROOT_DIR}/../api/docs/public/swagger.json"
INPUT_SPEC="${1:-${DEFAULT_INPUT_SPEC}}"
OUTPUT_SPEC="${ROOT_DIR}/openapi/public.swagger.json"
PUBLIC_API_HOST="${RIXL_PUBLIC_API_HOST:-api.rixl.com}"
PUBLIC_API_SCHEME="${RIXL_PUBLIC_API_SCHEME:-https}"

if ! command -v jq >/dev/null 2>&1; then
	echo "jq is required to prepare the public SDK spec" >&2
	exit 1
fi

if [[ ! -f "${INPUT_SPEC}" ]]; then
	echo "public swagger spec not found at ${INPUT_SPEC}" >&2
	echo "pass the source swagger.json path explicitly if api repo is not checked out next to this repo" >&2
	exit 1
fi

mkdir -p "$(dirname -- "${OUTPUT_SPEC}")"

jq --arg host "${PUBLIC_API_HOST}" --arg scheme "${PUBLIC_API_SCHEME}" '
	def strip_parameter_examples:
		walk(
			if type == "object" and has("in") then
				del(.example)
			else
				.
			end
		);

	def service_tags_catalog:
		[
			{
				"name": "Feeds",
				"description": "**Build dynamic social experiences.**<br><br>Feeds are ordered collections of Posts, designed for creating vertical video carousels, stories, or social timelines. Each Post contains your media (video or image) and its associated metadata."
			},
			{
				"name": "Videos",
				"description": "**Upload once, stream perfectly anywhere.**<br><br>RIXL's intelligent processing pipeline analyzes your video and transcodes it for optimal quality and performance. We handle the complexity so you can deliver a perfect playback experience every time."
			},
			{
				"name": "Images",
				"description": "**Deliver fast, perfectly formatted images.**<br><br>RIXL automatically handles image optimization, format conversion, and ThumbHash generation. All images are served via our global CDN for an ideal user experience."
			}
		];

	def sanitize_operation_id:
		gsub("[^A-Za-z0-9_]"; "_")
		| gsub("_+"; "_")
		| sub("^_"; "")
		| sub("_$"; "");

	def normalized_tag_token($tag):
		($tag // "")
		| ascii_downcase
		| gsub("[^a-z0-9]+"; "");

	def normalize_tag_label($tag):
		($tag // "")
		| gsub("[^A-Za-z0-9 ]"; " ")
		| split(" ")
		| map(select(length > 0))
		| map((.[0:1] | ascii_upcase) + .[1:])
		| join("");

	def canonical_service_tag($tag):
		(normalized_tag_token($tag)) as $token
		| if (["feed", "feeds", "post", "posts", "timeline", "timelines"] | index($token)) != null then
			"Feeds"
		elif (["video", "videos", "subtitle", "subtitles", "audiotrack", "audiotracks"] | index($token)) != null then
			"Videos"
		elif (["image", "images", "thumbnail", "thumbnails"] | index($token)) != null then
			"Images"
		else
			""
		end;

	def infer_service_tag($path):
		if $path | test("^/feeds(/|$)") then
			"Feeds"
		elif $path | test("^/videos(/|$)") then
			"Videos"
		elif $path | test("^/images(/|$)") then
			"Images"
		else
			""
		end;

	def default_operation_id($method; $path):
		($method + "_" + ($path
			| gsub("\\{"; "")
			| gsub("\\}"; "")
			| gsub("/"; "_")
			| gsub("-"; "_")
			| sanitize_operation_id));

	def is_http_method($method):
		(["get", "post", "put", "patch", "delete", "options", "head"] | index($method)) != null;

	def normalize_service_tags:
		.paths |= (
			to_entries
			| map(
				.key as $path
				| .value |= (
					to_entries
					| map(
						.key as $method
						| if is_http_method($method) then
							.value |= (
								(.tags // []) as $original_tags
								| .tags = [
									(
										(($original_tags | map(canonical_service_tag(.)) | map(select(length > 0)) | .[0])
											// infer_service_tag($path)
											// normalize_tag_label($original_tags[0])
											// "General")
									)
								]
							)
						else
							.
						end
					)
					| from_entries
				)
			)
			| from_entries
		);

	def rebuild_top_level_tags:
		(
			[
				.paths
				| to_entries[]
				| .value
				| to_entries[]
				| select(is_http_method(.key))
				| .value.tags[0]
			]
			| map(select(. != null and . != ""))
			| unique
		) as $used_tags
		| .tags = (
			if ($used_tags | length) > 0 then
				$used_tags
				| map(
					. as $name
					| (service_tags_catalog[] | select(.name == $name)) // {"name": $name, "description": ""}
				)
			else
				service_tags_catalog
			end
		);

	def ensure_operation_ids:
		.paths |= (
			to_entries
			| map(
				.key as $path
				| .value |= (
					to_entries
					| map(
						.key as $method
						| if is_http_method($method) then
							.value |= if ((.operationId // "") | length) == 0
								then . + {operationId: default_operation_id($method; $path)}
								else .
							end
						else
							.
						end
					)
					| from_entries
				)
			)
			| from_entries
		);

	def normalize_multipart_file_parameters:
		walk(
			if type == "object" and has("in") and .in == "formData" then
				if ((.type // "") == "file") then
					. + {type: "string", format: "binary"}
				elif ((.type // "") == "array") and (.items? != null) and ((.items.type // "") == "file") then
					.items = (.items + {type: "string", format: "binary"})
				else
					.
				end
			else
				.
			end
		);

	def strip_bearer_security:
		walk(
			if type == "object" and has("security") and (.security | type) == "array" then
				.security = [ .security[] | select(has("Bearer") | not) ]
				| if (.security | length) == 0 then del(.security) else . end
			else
				.
			end
		);

	strip_parameter_examples
	| ensure_operation_ids
	| normalize_service_tags
	| rebuild_top_level_tags
	| normalize_multipart_file_parameters
	| strip_bearer_security
	| if has("securityDefinitions") then .securityDefinitions |= del(.Bearer) else . end
	| .host = $host
	| .schemes = [$scheme]
	| .basePath = "/"
' "${INPUT_SPEC}" > "${OUTPUT_SPEC}"

if ! jq -e '
	[
		.paths
		| to_entries[]
		| .value
		| to_entries[]
		| select((["get", "post", "put", "patch", "delete", "options", "head"] | index(.key)) != null)
		| ((.value.tags // []) | length == 1)
	]
	| all
' "${OUTPUT_SPEC}" >/dev/null; then
	echo "prepared spec validation failed: every HTTP operation must have exactly one service tag" >&2
	exit 1
fi

echo "Prepared SDK input spec at ${OUTPUT_SPEC}"
