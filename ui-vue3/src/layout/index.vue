<!--
  ~ Licensed to the Apache Software Foundation (ASF) under one or more
  ~ contributor license agreements.  See the NOTICE file distributed with
  ~ this work for additional information regarding copyright ownership.
  ~ The ASF licenses this file to You under the Apache License, Version 2.0
  ~ (the "License"); you may not use this file except in compliance with
  ~ the License.  You may obtain a copy of the License at
  ~
  ~     http://www.apache.org/licenses/LICENSE-2.0
  ~
  ~ Unless required by applicable law or agreed to in writing, software
  ~ distributed under the License is distributed on an "AS IS" BASIS,
  ~ WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
  ~ See the License for the specific language governing permissions and
  ~ limitations under the License.
-->
<template>
  <div class="__container_layout_index">
    <a-layout style="min-height: 100vh">
      <a-layout-sider
        width="268"
        v-model:collapsed="collapsed"
        theme="light"
        :trigger="null"
        collapsible
      >
        <div class="logo">
          <img :src="logo" />
          <template v-if="!collapsed">Dubbo Admin</template>
        </div>
        <layout-menu></layout-menu>
      </a-layout-sider>
      <a-layout>
        <layout_header :collapsed="collapsed"></layout_header>
        <layout_bread></layout_bread>
        <a-layout-content
          :style="{ margin: '16px', padding: '16px', background: '#fff', minHeight: '280px' }"
        >
          <router-view />
        </a-layout-content>
        <a-layout-footer>todo</a-layout-footer>
      </a-layout>
    </a-layout>
  </div>
</template>
<script lang="ts" setup>
import { provide, ref } from 'vue'
import layoutMenu from './menu/layout_menu.vue'
import logo from '@/assets/logo.png'
import Layout_header from '@/layout/header/layout_header.vue'
import { PROVIDE_INJECT_KEY } from '@/base/enums/ProvideInject'
import Layout_bread from '@/layout/breadcrumb/layout_bread.vue'
import { PRIMARY_COLOR } from '@/base/constants'

let __null = PRIMARY_COLOR
const collapsed = ref<boolean>(false)
provide(PROVIDE_INJECT_KEY.COLLAPSED, collapsed)
</script>
<style lang="less" scoped>
.__container_layout_index {
  .logo {
    height: 40px;
    width: auto;
    margin: 10px 15px;
    padding-left: 10px;
    padding-right: 10px;
    border-radius: 8px;
    background: v-bind('PRIMARY_COLOR');
    line-height: 40px;
    vertical-align: middle;
    font-size: 22px;
    color: white;

    img {
      width: 28px;
      height: 28px;
      margin-bottom: 5px;
      margin-right: 5px;
    }
  }
}
</style>
