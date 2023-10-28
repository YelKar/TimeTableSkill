package main

import (
	"fmt"
)

func main() {
	resp, _ := Handler(HTTPRequest{Body: "{\"meta\":{\"locale\":\"ru-RU\",\"timezone\":\"UTC\",\"client_id\":\"ru.yandex.searchplugin/7.16 (none none; android 4.4.2)\",\"interfaces\":{\"screen\":{},\"payments\":{},\"account_linking\":{}}},\"session\":{\"message_id\":0,\"session_id\":\"7b909f4e-6b49-4ec1-b0cb-812a806b0a6f\",\"skill_id\":\"af6bf25c-0ae7-4910-b812-ff5491f4c63c\",\"user\":{\"user_id\":\"1A0742E8837DC797675A2CB4AD5C86D2414BD528E03315266E4E5F6891ADA17D\"},\"application\":{\"application_id\":\"4096D7308C7F809E254C43199E15719E1831B53EF88AEA649EF570B819391EBD\"},\"new\":true,\"user_id\":\"4096D7308C7F809E254C43199E15719E1831B53EF88AEA649EF570B819391EBD\"},\"request\":{\"command\":\"завтра\",\"original_utterance\":\"asd\",\"nlu\":{\"tokens\":[\"asd\"],\"entities\":[],\"intents\":{}},\"markup\":{\"dangerous_context\":false},\"type\":\"SimpleUtterance\"},\"version\":\"1.0\"}"})
	fmt.Println(resp.Body.Response.Text)
}
