// Copyright Project Harbor Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package rds

import (
	"fmt"

	"github.com/gomodule/redigo/redis"
)

const (
	requeueKeysPerJob = 4
)

// luaFuncStCodeText is common lua script function
var luaFuncStCodeText = `
-- for easily compare status
local stMap = { ['Pending'] = 0, ['Scheduled'] = 1, ['Running'] = 2, ['Success'] = 3, ['Stopped'] = 3, ['Error'] = 3 }

local function stCode(status)
  -- return 0 as default status
  return stMap[status] or 0
end
`

// luaFuncCompareText is common lua script function
var luaFuncCompareText = `
local function compare(status, revision)
  local sCode = stCode(status)
  local aCode = stCode(ARGV[1])
  local aRev = tonumber(ARGV[2]) or 0
  local aCheckInT = tonumber(ARGV[3]) or 0
  if revision < aRev or 
    ( revision == aRev and sCode <= aCode ) or
    ( revision == aRev and aCheckInT ~= 0 )
  then
     return 'ok'
  end
  return 'no'
end
`

// Script used to set the status of the job
//
// KEY[1]: key of job stats
// KEY[2]: key of inprogress track
// ARGV[1]: status text
// ARGV[2]: stats revision
// ARGV[3]: update timestamp
// ARGV[4]: job ID
var setStatusScriptText = fmt.Sprintf(`
%s

local res, st, code
`, luaFuncStCodeText)

// SetStatusScript is lua script for setting job status atomically
var SetStatusScript = redis.NewScript(2, setStatusScriptText)

var hookAckScriptText = fmt.Sprintf(`
%s

%s


`, luaFuncStCodeText, luaFuncCompareText)

// HookAckScript is defined to set the hook event ACK in the job stats map
var HookAckScript = redis.NewScript(2, hookAckScriptText)
