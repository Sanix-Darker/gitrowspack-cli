package main

type pingResponse struct {
	Version  string `json:"version"`
	Message  string `json:"message"`
	ServerIP string `json:"server-ip"`
}
