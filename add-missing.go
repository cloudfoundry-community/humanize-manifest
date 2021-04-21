package main

import (
	"fmt"
	"os"
	"strconv"

	"gopkg.in/yaml.v2"
)

type StoreMapSlice func(yaml.MapSlice)

// Adds to 'dest' the nodes that might found in 'source' and not 'dest'. Both
// source and dest are ordered YAML trees, so we deal with yaml.MapSlice types
// instead of map[interface{}]interface{}.
//
// Important assumption: dest YAML tree is a subset of the source YAML tree
func appendMissingNodes(source interface{}, dest interface{}, storeDest StoreMapSlice, path string) error {
	if debug {
		fmt.Fprintf(os.Stderr, "DEBUG: visiting path '%s' of type '%T'/'%T'\n", path, source, dest)
	}
	switch typedSrc := source.(type) {
	case yaml.MapSlice:
		typedDest := dest.(yaml.MapSlice)
		if debug {
			fmt.Fprintf(os.Stderr, "DEBUG: found MapItem slice at path '%s'\n", path)
		}
		for _, srcItem := range typedSrc {
			key := srcItem.Key.(string)
			srcValue := srcItem.Value

			dstItem, exists := findItemWithKey(key, typedDest)
			if !exists {
				if debug {
					fmt.Fprintf(os.Stderr, "DEBUG: adding missing value at path '%s'\n", path+"/"+key)
				}
				typedDest = append(typedDest, yaml.MapItem{Key: key, Value: srcValue})
				storeDest(typedDest)
				continue
			}
			storeChild := func(modified yaml.MapSlice) {
				if debug {
					fmt.Fprintf(os.Stderr, "DEBUG: storing modified MapSlice at path '%s' (MapSlice)\n", path+"/"+key)
				}
				(*dstItem).Value = modified
				storeDest(typedDest)
			}
			err := appendMissingNodes(srcValue, (*dstItem).Value, storeChild, path+"/"+key)
			if err != nil {
				return err
			}
		}

	case []interface{}:
		typedDest := dest.([]interface{})
		if len(typedSrc) != len(typedDest) {
			return fmt.Errorf("unexpected situation with arrays having different size at path '%s'.", path)
		}
		for idx := range typedSrc {
			storeChild := func(modified yaml.MapSlice) {
				if debug {
					fmt.Fprintf(os.Stderr, "DEBUG: storing modified MapSlice at path '%s' (slice)\n", path+"/"+strconv.Itoa(idx))
				}
				typedDest[idx] = modified
			}
			err := appendMissingNodes(typedSrc[idx], typedDest[idx], storeChild, path+"/"+strconv.Itoa(idx))
			if err != nil {
				return err
			}
		}
	case map[interface{}]interface{}:
		return fmt.Errorf("unexpected type 'map[interface{}]interface{}' in YAML tree, at path '%s'.", path)
		// typedDest := dest.(map[interface{}]interface{})
		// for untypedKey, srcValue := range typedSrc {
		// 	key := untypedKey.(string)
		// 	dstValue, exists := typedDest[key]
		// 	if !exists {
		// 		fmt.Printf("DEBUG: adding missing value at path '%s'\n", path+"/"+key)
		// 		typedDest[key] = srcValue
		// 		continue
		// 	}
		// 	storeChild := func(modified yaml.MapSlice) {
		// 		fmt.Printf("DEBUG: storing modified MapSlice at path '%s' (map)\n", path+"/"+key)
		// 		typedDest[key] = modified
		// 	}
		// 	err := appendMissingNodes(srcValue, dstValue, storeChild, path+"/"+key)
		// 	if err != nil {
		// 		return err
		// 	}
		// }
	}
	return nil
}

func findItemWithKey(key string, slice yaml.MapSlice) (*yaml.MapItem, bool) {
	var foundItem *yaml.MapItem
	exists := false
	for idx := range slice {
		if slice[idx].Key.(string) == key {
			exists = true
			foundItem = &slice[idx]
			break
		}
	}
	return foundItem, exists
}
