/*
 * Copyright 2025 CloudWeGo Authors
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

package milvus

import (
	"context"
	"encoding/binary"
	"math"

	"github.com/cloudwego/eino/callbacks"
	"github.com/cloudwego/eino/components"
	"github.com/cloudwego/eino/components/embedding"
)

// vector2Bytes converts vector to bytes
func vector2Bytes(vector []float64) []byte {
	float32Arr := make([]float32, len(vector))
	for i, v := range vector {
		float32Arr[i] = float32(v)
	}
	bytes := make([]byte, len(float32Arr)*4)
	for i, v := range float32Arr {
		binary.LittleEndian.PutUint32(bytes[i*4:], math.Float32bits(v))
	}
	return bytes
}

// MakeEmbeddingCtx makes the embedding context.
func makeEmbeddingCtx(ctx context.Context, emb embedding.Embedder) context.Context {
	runInfo := &callbacks.RunInfo{
		Component: components.ComponentOfEmbedding,
	}

	if embType, ok := components.GetType(emb); ok {
		runInfo.Type = embType
	}

	runInfo.Name = runInfo.Type + string(runInfo.Component)

	return callbacks.ReuseHandlers(ctx, runInfo)
}
