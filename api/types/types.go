// Copyright 2018 Globo.com authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package types

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

// Repository is the struct that stores all data from repository to be analyzed.
type Repository struct {
	ID             bson.ObjectId `bson:"_id,omitempty"`
	URL            string        `bson:"repositoryURL" json:"repositoryURL"`
	Branch         string        `json:"repositoryBranch"`
	CreatedAt      time.Time     `bson:"createdAt" json:"createdAt"`
	InternalDepURL string        `json:"internaldepURL"`
}

// SecurityTest is the struct that stores all data from the security tests to be executed.
type SecurityTest struct {
	ID               bson.ObjectId `bson:"_id,omitempty"`
	Name             string        `bson:"name" json:"name"`
	Image            string        `bson:"image" json:"image"`
	Cmd              string        `bson:"cmd" json:"cmd"`
	Language         string        `bson:"language" json:"language"`
	Default          bool          `bson:"default" json:"default"`
	TimeOutInSeconds int           `bson:"timeOutSeconds" json:"timeOutSeconds"`
}

// Analysis is the struct that stores all data from analysis performed.
type Analysis struct {
	ID             bson.ObjectId  `bson:"_id,omitempty"`
	RID            string         `bson:"RID" json:"RID"`
	URL            string         `bson:"repositoryURL" json:"repositoryURL"`
	Branch         string         `bson:"repositoryBranch" json:"repositoryBranch"`
	SecurityTests  []SecurityTest `bson:"securityTests" json:"securityTests"`
	Status         string         `bson:"status" json:"status"`
	Result         string         `bson:"result" json:"result"`
	Containers     []Container    `bson:"containers" json:"containers"`
	StartedAt      time.Time      `bson:"startedAt" json:"startedAt"`
	FinishedAt     time.Time      `bson:"finishedAt" json:"finishedAt"`
	InternalDepURL string         `bson:"internaldepURL,omitempty" json:"internaldepURL"`
	Codes          []Code         `bson:"codes" json:"codes"`
}

// Container is the struct that stores all data from a container run.
type Container struct {
	CID          string       `bson:"CID" json:"CID"`
	SecurityTest SecurityTest `bson:"securityTest" json:"securityTest"`
	CStatus      string       `bson:"cStatus" json:"cStatus"`
	COuput       string       `bson:"cOutput" json:"cOutput"`
	CResult      string       `bson:"cResult" json:"cResult"`
	CInfo        string       `bson:"cInfo" json:"cInfo"`
	StartedAt    time.Time    `bson:"startedAt" json:"startedAt"`
	FinishedAt   time.Time    `bson:"finishedAt" json:"finishedAt"`
}

// Code is the struct that stores all data from code found in a repository.
type Code struct {
	Language string   `bson:"language" json:"language"`
	Files    []string `bson:"files" json:"files"`
}
