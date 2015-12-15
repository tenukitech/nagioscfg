package nagioscfg

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"log"
	"regexp"
	"strings"
)

func ParseString(config string) (results []GenericObject, err error) {
	return parse(bytes.NewBufferString(config))
}

func parse(buf io.Reader) (objects []GenericObject, err error) {
	scanner := bufio.NewScanner(buf)
	defineRE := regexp.MustCompile(`define\s+(\S+)\s*{`)
	includeFileRE := regexp.MustCompile(`include_file=(.*)`)
	includeDirRE := regexp.MustCompile(`include_dir=(.*)`)
	attributeRE := regexp.MustCompile(`(\S+)\s+(.+)`)
	var obj GenericObject

	for scanner.Scan() {
		line := strings.Trim(scanner.Text(), " \t")
		log.Println("Parsing: " + line)

		if strings.HasPrefix(line, "define") {

			match := defineRE.FindStringSubmatch(line)
			if match == nil {
				err = fmt.Errorf("Invalid configuration line '%s'", line)
				return
			}

			log.Printf("Creating new object with type %s", match[1])

			if obj.ObjectType != "" {
				objects = append(objects, obj)
			}

			obj = NewGenericObject(match[1])
			continue
		}

		if strings.HasPrefix(line, "#") || strings.HasPrefix(line, ";") {
			continue // Skip comments
		}

		if match := includeFileRE.FindStringSubmatch(line); match != nil {
			fmt.Printf("TODO: InludeFile")
			continue
		}

		if match := includeDirRE.FindStringSubmatch(line); match != nil {
			fmt.Printf("TODO: InludeDir")
			continue
		}

		if match := attributeRE.FindStringSubmatch(line); match != nil {
			obj.Attributes[match[1]] = match[2]
		}
	}

	if obj.ObjectType != "" {
		objects = append(objects, obj)
	}

	err = scanner.Err()

	return
}
