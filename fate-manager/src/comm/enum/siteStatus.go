/*
 * Copyright 2020 The FATE Authors. All Rights Reserved.
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
package enum

import "fate.manager/entity"

type SiteStatusType int32

const (
	SITE_STATUS_UNKNOWN SiteStatusType = -1
	SITE_STATUS_NO_JOIN SiteStatusType = 1
	SITE_STATUS_JOINED  SiteStatusType = 2
	SITE_STATUS_REMOVED SiteStatusType = 3
)

func GetSiteString(p SiteStatusType) string {
	switch p {
	case SITE_STATUS_NO_JOIN:
		return "Not Join"
	case SITE_STATUS_JOINED:
		return "Joined"
	case SITE_STATUS_REMOVED:
		return "Deleted"
	default:
		return "unknown"
	}
}

func GetSiteStatusList() []entity.IdPair {
	var idPairList []entity.IdPair
	for i := 0; i < 4; i++ {
		idPair := entity.IdPair{i, GetSiteString(SiteStatusType(i))}
		idPairList = append(idPairList, idPair)
	}
	return idPairList
}
