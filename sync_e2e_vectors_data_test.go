package opencode_test

// Code generated from openapi.json by .tmp/sync-1.18.15/gen_vectors.py.
// DO NOT EDIT by hand. These constants are the OpenAPI-derived legal JSON instances.

var syncVecConfigAutoupdate = map[string]string{
	"boolean_true":  "true",
	"boolean_false": "false",
	"notify":        "\"notify\"",
}

var syncVecConfigFormatter = map[string]string{
	"boolean_false": "false",
	"boolean_true":  "true",
	"map_single":    "{\"prettier\":{\"disabled\":true,\"command\":[\"sample-string\"],\"environment\":{},\"extensions\":[\"sample-string\"]}}",
	"map_multi":     "{\"prettier\":{\"disabled\":false,\"command\":[\"prettier\",\"--write\"],\"environment\":{\"NODE_ENV\":\"production\"},\"extensions\":[\".ts\"]},\"biome\":{\"disabled\":true}}",
}

var syncVecConfigLsp = map[string]string{
	"boolean_false":       "false",
	"boolean_true":        "true",
	"map_server_disabled": "{\"gopls\":{\"disabled\":true}}",
	"map_server_command":  "{\"gopls\":{\"command\":[\"sample-string\"],\"extensions\":[\"sample-string\"],\"disabled\":true,\"env\":{},\"initialization\":{}}}",
	"map_mixed":           "{\"gopls\":{\"command\":[\"gopls\"],\"extensions\":[\".go\"],\"env\":{\"GOPLS_GENERATE\":\"1\"},\"initialization\":{\"formatting\":true}},\"python\":{\"disabled\":true}}",
}

var syncVecConfigPlugin = map[string]string{
	"single_name":  "[\"sample-string\"]",
	"single_tuple": "[[\"sample-string\",{}]]",
	"mixed":        "[\"sample-string\",[\"sample-string\",{}]]",
}

var syncVecConfigReference = map[string]string{
	"string":    "{\"x\":\"sample-string\"}",
	"git":       "{\"x\":{\"repository\":\"sample-string\",\"branch\":\"sample-string\",\"description\":\"sample-string\",\"hidden\":true}}",
	"local":     "{\"x\":{\"path\":\"sample-string\",\"description\":\"sample-string\",\"hidden\":true}}",
	"map_mixed": "{\"plain\":\"https://example.com/docs\",\"git\":{\"repository\":\"owner/repo\",\"branch\":\"main\",\"description\":\"d\",\"hidden\":false},\"local\":{\"path\":\"./docs\",\"description\":\"d\",\"hidden\":true}}",
}

var syncVecProviderInterleaved = map[string]string{
	"boolean_true":           "true",
	"boolean_false":          "false",
	"enum_reasoning":         "\"reasoning\"",
	"enum_reasoning_content": "\"reasoning_content\"",
	"enum_reasoning_text":    "\"reasoning_text\"",
	"vendor_string":          "\"vendor_custom\"",
	"object_field":           "{\"field\":\"reasoning\"}",
}

var syncVecAPIError = map[string]string{
	"full":    "{\"name\":\"APIError\",\"data\":{\"message\":\"sample-string\",\"statusCode\":1,\"isRetryable\":true,\"responseHeaders\":{},\"responseBody\":\"sample-string\",\"metadata\":{}}}",
	"minimal": "{\"name\":\"APIError\",\"data\":{\"message\":\"boom\",\"isRetryable\":true}}",
}

var syncVecEventPtyExited = map[string]string{
	"exit_0":          "{\"id\":\"evt_1\",\"type\":\"pty.exited\",\"properties\":{\"id\":\"pty_1\",\"exitCode\":0}}",
	"exit_max_int32":  "{\"id\":\"evt_2\",\"type\":\"pty.exited\",\"properties\":{\"id\":\"pty_2\",\"exitCode\":2147483647}}",
	"exit_2147483648": "{\"id\":\"evt_3\",\"type\":\"pty.exited\",\"properties\":{\"id\":\"pty_3\",\"exitCode\":2147483648}}",
	"exit_max_int64":  "{\"id\":\"evt_4\",\"type\":\"pty.exited\",\"properties\":{\"id\":\"pty_4\",\"exitCode\":9007199254740991}}",
}

var syncVecSyncEvents = map[string]string{
	"context_updated": "{\"type\":\"sync\",\"id\":\"evt_sync\",\"syncEvent\":{\"type\":\"session.next.context.updated.1\",\"id\":\"sample-string\",\"seq\":1,\"aggregateID\":\"sample-string\",\"data\":{\"timestamp\":1,\"sessionID\":\"sample-string\",\"messageID\":\"sample-string\",\"text\":\"sample-string\"}}}",
	"prompt_admitted": "{\"type\":\"sync\",\"id\":\"evt_sync\",\"syncEvent\":{\"type\":\"session.next.prompt.admitted.1\",\"id\":\"sample-string\",\"seq\":1,\"aggregateID\":\"sample-string\",\"data\":{\"timestamp\":1,\"sessionID\":\"sample-string\",\"messageID\":\"sample-string\",\"prompt\":{\"text\":\"sample-string\",\"files\":[{\"uri\":\"sample-string\",\"mime\":\"sample-string\",\"name\":\"sample-string\",\"description\":\"sample-string\",\"source\":{\"start\":1,\"end\":1,\"text\":\"sample-string\"}}],\"agents\":[{\"name\":\"sample-string\",\"source\":{\"start\":1,\"end\":1,\"text\":\"sample-string\"}}]},\"delivery\":\"steer\"}}}",
}

var syncVecGlobalEvents = map[string]string{
	"member_0_models-dev.refreshed":          "{\"directory\":\"/repo\",\"project\":\"prj_1\",\"workspace\":\"ws_1\",\"payload\":{\"id\":\"sample-string\",\"type\":\"models-dev.refreshed\",\"properties\":{}}}",
	"member_4_session.created":               "{\"directory\":\"/repo\",\"project\":\"prj_1\",\"workspace\":\"ws_1\",\"payload\":{\"id\":\"sample-string\",\"type\":\"session.created\",\"properties\":{\"sessionID\":\"sample-string\",\"info\":{\"id\":\"sample-string\",\"slug\":\"sample-string\",\"projectID\":\"sample-string\",\"workspaceID\":\"sample-string\",\"directory\":\"sample-string\",\"path\":\"sample-string\",\"parentID\":\"sample-string\",\"summary\":{\"additions\":1,\"deletions\":1,\"files\":1,\"diffs\":[{\"file\":\"sample-string\",\"patch\":\"sample-string\",\"additions\":1,\"deletions\":1,\"status\":\"added\"}]},\"cost\":1,\"tokens\":{\"input\":1,\"output\":1,\"reasoning\":1,\"cache\":{\"read\":1,\"write\":1}},\"share\":{\"url\":\"sample-string\"},\"title\":\"sample-string\",\"agent\":\"sample-string\",\"model\":{\"id\":\"sample-string\",\"providerID\":\"sample-string\",\"variant\":\"sample-string\"},\"version\":\"sample-string\",\"metadata\":{},\"time\":{\"created\":1,\"updated\":1,\"compacting\":1,\"archived\":1},\"permission\":[{\"permission\":\"sample-string\",\"pattern\":\"sample-string\",\"action\":\"allow\"}],\"revert\":{\"messageID\":\"sample-string\",\"partID\":\"sample-string\",\"snapshot\":\"sample-string\",\"diff\":\"sample-string\"}}}}}",
	"member_9_message.part.updated":          "{\"directory\":\"/repo\",\"project\":\"prj_1\",\"workspace\":\"ws_1\",\"payload\":{\"id\":\"sample-string\",\"type\":\"message.part.updated\",\"properties\":{\"sessionID\":\"sample-string\",\"part\":{\"id\":\"sample-string\",\"sessionID\":\"sample-string\",\"messageID\":\"sample-string\",\"type\":\"text\",\"text\":\"sample-string\",\"synthetic\":true,\"ignored\":true,\"time\":{\"start\":1,\"end\":1},\"metadata\":{}},\"time\":1}}}",
	"member_15_session.next.prompt.admitted": "{\"directory\":\"/repo\",\"project\":\"prj_1\",\"workspace\":\"ws_1\",\"payload\":{\"id\":\"sample-string\",\"type\":\"session.next.prompt.admitted\",\"properties\":{\"timestamp\":1,\"sessionID\":\"sample-string\",\"messageID\":\"sample-string\",\"prompt\":{\"text\":\"sample-string\",\"files\":[{\"uri\":\"sample-string\",\"mime\":\"sample-string\",\"name\":\"sample-string\",\"description\":\"sample-string\",\"source\":{\"start\":1,\"end\":1,\"text\":\"sample-string\"}}],\"agents\":[{\"name\":\"sample-string\",\"source\":{\"start\":1,\"end\":1,\"text\":\"sample-string\"}}]},\"delivery\":\"steer\"}}}",
	"member_16_session.next.context.updated": "{\"directory\":\"/repo\",\"project\":\"prj_1\",\"workspace\":\"ws_1\",\"payload\":{\"id\":\"sample-string\",\"type\":\"session.next.context.updated\",\"properties\":{\"timestamp\":1,\"sessionID\":\"sample-string\",\"messageID\":\"sample-string\",\"text\":\"sample-string\"}}}",
	"member_44_session.diff":                 "{\"directory\":\"/repo\",\"project\":\"prj_1\",\"workspace\":\"ws_1\",\"payload\":{\"id\":\"sample-string\",\"type\":\"session.diff\",\"properties\":{\"sessionID\":\"sample-string\",\"diff\":[{\"file\":\"sample-string\",\"patch\":\"sample-string\",\"additions\":1,\"deletions\":1,\"status\":\"added\"}]}}}",
	"member_48_file.edited":                  "{\"directory\":\"/repo\",\"project\":\"prj_1\",\"workspace\":\"ws_1\",\"payload\":{\"id\":\"sample-string\",\"type\":\"file.edited\",\"properties\":{\"file\":\"sample-string\"}}}",
	"member_57_pty.exited":                   "{\"directory\":\"/repo\",\"project\":\"prj_1\",\"workspace\":\"ws_1\",\"payload\":{\"id\":\"sample-string\",\"type\":\"pty.exited\",\"properties\":{\"id\":\"sample-string\",\"exitCode\":1}}}",
	"member_79_session.compacted":            "{\"directory\":\"/repo\",\"project\":\"prj_1\",\"workspace\":\"ws_1\",\"payload\":{\"id\":\"sample-string\",\"type\":\"session.compacted\",\"properties\":{\"sessionID\":\"sample-string\"}}}",
	"member_89_SyncEventSessionCreated":      "{\"directory\":\"/repo\",\"project\":\"prj_1\",\"workspace\":\"ws_1\",\"payload\":{\"type\":\"sync\",\"id\":\"sample-string\",\"syncEvent\":{\"type\":\"session.created.1\",\"id\":\"sample-string\",\"seq\":1,\"aggregateID\":\"sample-string\",\"data\":{\"sessionID\":\"sample-string\",\"info\":{\"id\":\"sample-string\",\"slug\":\"sample-string\",\"projectID\":\"sample-string\",\"workspaceID\":\"sample-string\",\"directory\":\"sample-string\",\"path\":\"sample-string\",\"parentID\":\"sample-string\",\"summary\":{\"additions\":1,\"deletions\":1,\"files\":1,\"diffs\":[{\"file\":\"sample-string\",\"patch\":\"sample-string\",\"additions\":1,\"deletions\":1,\"status\":\"added\"}]},\"cost\":1,\"tokens\":{\"input\":1,\"output\":1,\"reasoning\":1,\"cache\":{\"read\":1,\"write\":1}},\"share\":{\"url\":\"sample-string\"},\"title\":\"sample-string\",\"agent\":\"sample-string\",\"model\":{\"id\":\"sample-string\",\"providerID\":\"sample-string\",\"variant\":\"sample-string\"},\"version\":\"sample-string\",\"metadata\":{},\"time\":{\"created\":1,\"updated\":1,\"compacting\":1,\"archived\":1},\"permission\":[{\"permission\":\"sample-string\",\"pattern\":\"sample-string\",\"action\":\"allow\"}],\"revert\":{\"messageID\":\"sample-string\",\"partID\":\"sample-string\",\"snapshot\":\"sample-string\",\"diff\":\"sample-string\"}}}}}}",
	"sync_wrapper":                           "{\"directory\":\"/repo\",\"project\":\"prj_1\",\"workspace\":\"ws_1\",\"payload\":{\"type\":\"sync\",\"id\":\"evt_sync\",\"syncEvent\":{\"type\":\"session.next.prompt.admitted.1\",\"id\":\"sample-string\",\"seq\":1,\"aggregateID\":\"sample-string\",\"data\":{\"timestamp\":1,\"sessionID\":\"sample-string\",\"messageID\":\"sample-string\",\"prompt\":{\"text\":\"sample-string\",\"files\":[{\"uri\":\"sample-string\",\"mime\":\"sample-string\",\"name\":\"sample-string\",\"description\":\"sample-string\",\"source\":{\"start\":1,\"end\":1,\"text\":\"sample-string\"}}],\"agents\":[{\"name\":\"sample-string\",\"source\":{\"start\":1,\"end\":1,\"text\":\"sample-string\"}}]},\"delivery\":\"steer\"}}}}",
}
