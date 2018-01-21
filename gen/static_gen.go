package main

import (
	"github.com/viant/endly"
	"github.com/viant/toolbox"
	"github.com/viant/toolbox/storage"
	"log"
	"path"
)

//main generate file with static content from meta, workflow, req folders so that they can be compiled into final binary
func main() {
	genDirectory := toolbox.CallerDirectory(3)
	parent := string(genDirectory[:len(genDirectory)-4])
	mappings := []*storage.StorageMapping{
		{
			SourceURL:      toolbox.FileSchema + path.Join(parent, "meta"),
			DestinationURI: path.Join(endly.EndlyNamespace, "meta"),
			TargetFile:     path.Join(parent, "static", "meta.go"),
			TargetPackage:  "static",
		},
		{
			SourceURL:      toolbox.FileSchema + path.Join(parent, "workflow"),
			DestinationURI: path.Join(endly.EndlyNamespace, "workflow"),
			TargetFile:     path.Join(parent, "static", "workflow.go"),
			TargetPackage:  "static",
		},
		{
			SourceURL:      toolbox.FileSchema + path.Join(parent, "req"),
			DestinationURI: path.Join(endly.EndlyNamespace, "req"),
			TargetFile:     path.Join(parent, "static", "req.go"),
			TargetPackage:  "static",
		},
		{
			SourceURL:      toolbox.FileSchema + path.Join(parent, "Version"),
			DestinationURI: path.Join(endly.EndlyNamespace, "Version"),
			TargetFile:     path.Join(parent, "static", "version.go"),
			TargetPackage:  "static",
		},
	}
	err := storage.GenerateStorageCode(mappings...)
	if err != nil {
		log.Fatal(err)
	}
}