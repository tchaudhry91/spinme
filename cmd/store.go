/*
Copyright © 2019 Tanmay Chaudhry <tanmay.chaudhry@gmail.com>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"encoding/json"

	"github.com/boltdb/bolt"
	"github.com/tchaudhry91/spinme/spin"
)

func storeConfig(dbFile string, o spin.SpinOut) error {
	db, err := bolt.Open(dbFile, 0600, nil)
	if err != nil {
		return err
	}
	defer db.Close()
	err = db.Update(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists([]byte("spinners"))
		if err != nil {
			return err
		}
		enc, err := json.Marshal(o)
		if err != nil {
			return err
		}
		return b.Put([]byte(o.ID), enc)
	})
	return err
}

func removeConfig(dbFile string, id string) error {
	db, err := bolt.Open(dbFile, 0600, nil)
	if err != nil {
		return err
	}
	defer db.Close()
	err = db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("spinners"))
		return b.Delete([]byte(id))
	})
	return err
}

func getConfigs(dbFile string) ([]spin.SpinOut, error) {
	var oo []spin.SpinOut
	db, err := bolt.Open(dbFile, 0600, nil)
	if err != nil {
		return oo, err
	}
	defer db.Close()
	err = db.View(func(tx *bolt.Tx) error {
		if err != nil {
			return err
		}
		b := tx.Bucket([]byte("spinners"))
		if b == nil {
			return nil
		}
		err = b.ForEach(func(k []byte, v []byte) error {
			var o spin.SpinOut
			err := json.Unmarshal(v, &o)
			if err != nil {
				return err
			}
			oo = append(oo, o)
			return nil
		})
		if err != nil {
			return err
		}
		return nil
	})
	return oo, err
}
