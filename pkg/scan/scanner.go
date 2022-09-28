/*
 * Copyright 2022 zpxio (Jeff Sharpe)
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package scan

import (
	"github.com/apex/log"
	"os"
	"path/filepath"
)

type Scanner struct {
	rootPath    string
	scanTargets []string

	pruneTargets map[string]bool
	scanned      map[string]bool
}

func CreateScanner(dirs ...string) *Scanner {
	dirsClean := make([]string, len(dirs))
	for i, d := range dirs {
		dirsClean[i] = filepath.Clean(d)
	}

	return &Scanner{
		scanTargets: dirsClean,

		pruneTargets: map[string]bool{},
		scanned:      map[string]bool{},
	}
}

func (s *Scanner) ScanTo() error {
	var err error

	for _, d := range s.scanTargets {
		err = filepath.Walk(d,
			func(path string, info os.FileInfo, itemError error) error {

				if itemError != nil {
					log.Infof("Cannot scan item (%s): %s", path, itemError)
					return nil
				} else if info.IsDir() {
					if _, found := s.scanned[path]; found {
						log.Infof("Skipping previously scanned directory: %s", path)
						return filepath.SkipDir
					} else {
						s.scanned[path] = true
						s.scanDirectory(path, info)
					}
				} else { // Item is a file
					if info.Mode().Type() == os.ModeSymlink {
						s.scanSymlink(path, info)
					} else {
						s.scanFile(path, info)
					}
				}

				return nil
			})
	}

	return err
}

func (s *Scanner) scanFile(path string, info os.FileInfo) {
	log.Infof("Examining file: %s", path)
}

func (s *Scanner) scanDirectory(path string, info os.FileInfo) {
	log.Infof("Examining directory: %s", path)
}

func (s *Scanner) scanSymlink(path string, info os.FileInfo) {
	log.Infof("Examining symlink: %s", path)
}
