/*
 * Copyright (c) 2020. Temple3x (temple3x@gmail.com)
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

// Package settings is the global settings of zai.
// Don't modify it unless you totally know what will happen.
package settings

import (
	"time"

	"github.com/templexxx/crank/typeutil"
)

const (
	kb = 1024
	mb = 1024 * kb
	gb = 1024 * mb
)

const (
	// DefaultLogRoot is the default log files path root.
	// e.g.:
	// <DefaultLogRoot>/<appName>/access.log
	// & <DefaultLogRoot>/<appName>/error.log
	DefaultLogRoot = "/var/log/zai"
)

const (
	MaxObjectSize = 4 * mb
)

const (
	ExtReserved uint16 = iota
	ExtV1
	ExtV2
	ExtV3
	ExtV4
	ExtV5
)

var ValidExtVersions = []uint16{ExtV1, ExtV2, ExtV3, ExtV4, ExtV5}

// Zai has three different isolation levels.
const (
	IsolationInstance = "instance"
	IsolationDisk     = "disk"
	IsolationNone     = "none"
)

var ValidIsolationLevels = []string{IsolationInstance, IsolationDisk, IsolationNone}

// DefaultIsolationLevel is IsolationInstance, enough for giving enough protection:
// 1. Each machine in the same box will only be placed in the same IDC. (see arch docs for details)
// 2. Instance isolation is enough storage for giving high durability.
const DefaultIsolationLevel = IsolationInstance

// DefaultReplicas is 2.
// In Tesamc, we will start at 2 replicas first for saving overhead.
// There are other replicas in public/private cloud storage. 2 replicas just for speeding up repairing.
const DefaultReplicas = 2

const (
	ExtV1SegCnt = 256
)

const (
	// DefaultZBufHeartbeatInterval is the interval of two zBuf heartbeat in zBuf server.
	// Our env is stable, we won't too frequently heartbeat.
	DefaultZBufHeartbeatInterval = 15 * time.Second

	DefaultExtHeartbeatInterval = 30 * time.Second
)

// Different sizes of objects will be put into different versions of extents for avoiding high hash collision.
// In these settings, we will get 3% - 11.8% hash collision at most when extents are full.
// For details, see issue: https://g.tesamc.com/IT/zai/issues/40
const (
	DefaultV1SegmentSize = typeutil.ByteSize(4 * 1024 * 1024) // 4MB, Low hash collision for (0, 64KB] objects.
	DefaultV2SegmentSize = DefaultV1SegmentSize * 4           // 16MB,  Low hash collision for (64KB, 256KB] objects.
	DefaultV3SegmentSize = DefaultV2SegmentSize * 4           // 64MB, Low hash collision for (256KB, 1MB] objects.
	DefaultV4SegmentSize = DefaultV3SegmentSize * 4           // 256MB, Low hash collision for (1MB, 4MB) objects.
	DefaultV5SegmentSize = DefaultV4SegmentSize * 2           // 512MB, Low hash collision for [4MB] objects.
)
