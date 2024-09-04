package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", defaultHandler)

	fmt.Println("Server is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func defaultHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		fmt.Fprint(w, cv)
	default:
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}

const cv = `
  ----
  Info:
    name: Camin McCluskey
    email: camin(at)stackfix.com
    bio: I'm the Co-Founder & CTO of Stackfix, where we’re building an AI-powered marketplace that enables businesses to 
        buy, implement and optimise the right software in seconds. Previously, I co-founded another startup, 
        Telescope (we used ML for sales outreach) and was a software engineer at Skyscanner where I worked on identity 
        infrastructure.
  ----
  Links:
    github: github.com/camin-mccluskey
    linkedIn: linkedin.com/in/camin-mccluskey/
    x (twitter): x.com/caminmc
    personal site: camin.xyz
    blog: blog.camin.xyz
  ----
  Experience:
    Stackfix (2023 - Present):
      role: Co-Founder & CTO
    Telescope (2021 - 2023):
      role: Co-Founder & CTO
    Skyscanner (2019 - 2021):
      role: Software Engineer
  ----
  Skills & Interests:
    React, Next.js, Node.js, Python, Typescript, Golang, Clojure,
    Writing (blogging), AI, ML, SaaS, Startups
    Golf, Rugby
`
