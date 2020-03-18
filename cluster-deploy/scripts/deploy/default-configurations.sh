#!/bin/bash
#
# Copyright 2020 The FATE Authors. All Rights Reserved.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.
#

system_user=app
install_dir=
java_dir=
#cloud-manager
cloud_ip=
cloud_port=9999
cloud_db_ip=
cloud_db_name=cloud_manager
cloud_db_user=
cloud_db_password=
cloud_db_dir=
#fateboard
fateboard_branch=develop-1.2.1
fateboard_db_dir=
fateboard_db_ip=()
fateboard_db_name=
fateboard_db_user=
fateboard_db_password=
fateboard_ip=()
fateboard_port=8080
fate_flow_ip=()
fate_flow_port=9380
fate_path=