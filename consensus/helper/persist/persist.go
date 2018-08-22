/*
MIT License

Copyright (c) 2018 invin

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
*/

package persist

import (
	"github.com/hyperledger/fabric/core/db"
)

// Helper provides an abstraction to access the Persist column family
// in the database.
type Helper struct{}

// StoreState stores a key,value pair
func (h *Helper) StoreState(key string, value []byte) error {
	db := db.GetDBHandle()
	return db.Put(db.PersistCF, []byte("consensus."+key), value)
}

// DelState removes a key,value pair
func (h *Helper) DelState(key string) {
	db := db.GetDBHandle()
	db.Delete(db.PersistCF, []byte("consensus."+key))
}

// ReadState retrieves a value to a key
func (h *Helper) ReadState(key string) ([]byte, error) {
	db := db.GetDBHandle()
	return db.Get(db.PersistCF, []byte("consensus."+key))
}

// ReadStateSet retrieves all key,value pairs where the key starts with prefix
func (h *Helper) ReadStateSet(prefix string) (map[string][]byte, error) {
	db := db.GetDBHandle()
	prefixRaw := []byte("consensus." + prefix)

	ret := make(map[string][]byte)
	it := db.GetIterator(db.PersistCF)
	defer it.Close()
	for it.Seek(prefixRaw); it.ValidForPrefix(prefixRaw); it.Next() {
		key := string(it.Key().Data())
		key = key[len("consensus."):]
		// copy data from the slice!
		ret[key] = append([]byte(nil), it.Value().Data()...)
	}
	return ret, nil
}
