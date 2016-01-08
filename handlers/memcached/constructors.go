// Copyright 2015 Netflix, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package memcached

import "io"

import "github.com/netflix/rend/handlers/memcached/std"
import "github.com/netflix/rend/handlers/memcached/chunked"

func NewHandler(conn io.ReadWriteCloser) std.Handler {
	return std.NewHandler(conn)
}

func NewChunkedHandler(conn io.ReadWriteCloser) chunked.Handler {
	return chunked.NewHandler(conn)
}