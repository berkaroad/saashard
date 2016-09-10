// The MIT License (MIT)

// Copyright (c) 2016 Jerry Bai

// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:

// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.

// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

package route

import (
	"fmt"
	"strings"

	"github.com/berkaroad/saashard/sqlparser"
)

var hintPrefix = "/*!saashard "
var hintNodesPrefix = "nodes="

// Hint to extra process.
// Nodes: /*!saashard nodes=node1,node2 */
// OnMaster: /*!saashard master */
type Hint struct {
	OnMaster bool
	Nodes    []string
}

// ReadHint read hint from comments
func ReadHint(comments *sqlparser.Comments) *Hint {
	hint := new(Hint)
	for i, comment := range *comments {
		commentStr := string(comment)
		if strings.HasPrefix(commentStr, hintPrefix) {
			commentStr = strings.TrimSuffix(strings.TrimPrefix(commentStr, hintPrefix), "*/")
			commentStr = strings.ToLower(strings.TrimSpace(commentStr))
			if commentStr == "master" {
				hint.OnMaster = true
			} else if strings.HasPrefix(commentStr, hintNodesPrefix) {
				nodesStr := strings.TrimPrefix(commentStr, hintNodesPrefix)
				for _, nodeStr := range strings.Split(nodesStr, ",") {
					node := strings.TrimSpace(nodeStr)
					if !sqlparser.Contains(hint.Nodes, node) {
						hint.Nodes = append(hint.Nodes, node)
					}
				}
			}
			([][]byte)(*comments)[i] = []byte("")
		}
	}
	fmt.Printf("Hints.Nodes='%s'; Hints.OnMaster='%v'\r\n", strings.Join(hint.Nodes, ","), hint.OnMaster)
	return hint
}
