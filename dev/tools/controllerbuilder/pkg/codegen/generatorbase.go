// Copyright 2024 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package codegen

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"k8s.io/klog/v2"
)

type generatorBase struct {
	outputBaseDir  string
	generatedFiles map[generatedFileKey]*generatedFile
	errors         []error
}

func (g *generatorBase) init(outputBaseDir string) {
	g.outputBaseDir = outputBaseDir
	g.generatedFiles = make(map[generatedFileKey]*generatedFile)
}

func (g *generatorBase) getOutputFile(k generatedFileKey) *generatedFile {
	out := g.generatedFiles[k]
	if out == nil {
		out = &generatedFile{key: k, baseDir: g.outputBaseDir}
		g.generatedFiles[k] = out
	}
	return out
}

func (g *generatorBase) Errorf(msg string, args ...any) {
	g.errors = append(g.errors, fmt.Errorf(msg, args...))
}

type generatedFile struct {
	baseDir  string
	key      generatedFileKey
	contents bytes.Buffer
}

type generatedFileKey struct {
	GoPackage string

	FileName string
}

func (f *generatedFile) OutputDir() string {
	tokens := strings.Split(f.key.GoPackage, ".")
	dirTokens := []string{f.baseDir}
	dirTokens = append(dirTokens, tokens...)
	dir := filepath.Join(dirTokens...)
	return dir
}

func (f *generatedFile) Write(addCopyright bool) error {
	if f.contents.Len() == 0 {
		return nil
	}

	dir := f.OutputDir()

	if err := os.MkdirAll(dir, 0755); err != nil {
		return fmt.Errorf("creating directory %q: %w", dir, err)
	}

	var w bytes.Buffer

	if addCopyright {
		writeCopyright(&w, time.Now().Year())
	}

	f.contents.WriteTo(&w)

	p := filepath.Join(dir, f.key.FileName)
	klog.Infof("writing file %v", p)
	if err := os.WriteFile(p, w.Bytes(), 0644); err != nil {
		return fmt.Errorf("writing %q: %w", p, err)
	}

	return nil
}

func (v *generatorBase) WriteFiles(addCopyright bool) error {
	for _, f := range v.generatedFiles {
		if err := f.Write(addCopyright); err != nil {
			return err
		}
	}
	return nil
}
