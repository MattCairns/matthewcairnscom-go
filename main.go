package main

import (
	"html/template"
	"net/http"
)

type project struct {
	Screenshot  string
	Title       string
	Description string
	Source      string
}

var (
	tpl = template.Must(template.ParseGlob("templates/*"))
)

func main() {
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))
	http.HandleFunc("/", index)
	http.HandleFunc("/projects/", projects)
	http.ListenAndServe(":80", nil)
}

func projects(w http.ResponseWriter, req *http.Request) {
	headfoot(w, req)

	mandelbrot := project{
		Screenshot:  "https://d2wyztt7kkorr5.cloudfront.net/projects/mandelbrot.png",
		Title:       "mandelbrot set - threejs",
		Description: "an implementation of the mandelbrot set using shaders and threejs.",
		Source:      "https://github.com/mattcairns/source",
	}

	comiccon := project{
		Screenshot:  "https://d2wyztt7kkorr5.cloudfront.net/projects/comicConMenu.png",
		Title:       "the comic con - game",
		Description: "a group orject create for csc167. this is a 2d game made with gamemaker and gamemaker language, you play a comic book character who must escape the book through a series of puzzles.",
		Source:      "https://github.com/mattcairns/thecomiccon",
	}

	voxelgen := project{
		Screenshot:  "https://d2wyztt7kkorr5.cloudfront.net/projects/VoxelTerrainScreen.png",
		Title:       "voxel terrain generator",
		Description: "create a randomly generated voxel terrain of a specified size. written in java using the lwjgl library.",
		Source:      "https://github.com/mattcairns/voxelterraingenerator",
	}

	pong := project{
		Screenshot:  "https://d2wyztt7kkorr5.cloudfront.net/projects/pongScreen.png",
		Title:       "Classic Pong Remake",
		Description: "A remake of the classic game of pong written in Java using the LibGDX library.",
		Source:      "https://github.com/mattcairns/pong",
	}

	huntsarah := project{
		Screenshot:  "https://d2wyztt7kkorr5.cloudfront.net/projects/huntSarahScreen.png",
		Title:       "Hunt the Sarah",
		Description: "Based on the classic Gauntlet game you must hunt the Sarahs and destroy their spawners.  Written using the LibGDX library.",
		Source:      "https://github.com/mattcairns/HuntTheSarahPrototype",
	}

	projects := []project{mandelbrot, comiccon, voxelgen, pong, huntsarah}

	tpl.ExecuteTemplate(w, "projects.gohtml", projects)

}

func index(w http.ResponseWriter, req *http.Request) {
	headfoot(w, req)
	images := [5]string{"mountains", "cloudExpload", "MRC_1", "MRC_2", "MRC_3"}
	tpl.ExecuteTemplate(w, "index.gohtml", images)
}

func headfoot(w http.ResponseWriter, req *http.Request) {
	tpl.ExecuteTemplate(w, "header.gohtml", nil)
	tpl.ExecuteTemplate(w, "footer.gohtml", nil)

}
