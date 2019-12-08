<template>
  <v-app>
    <v-app-bar
      app
      dense
      dark
      :height="56"
    >
      <v-toolbar-title>
        <router-link
          class="app__title"
          :to="{ name: routerToHomeName }"
        >
          test_auth_app
        </router-link>
      </v-toolbar-title>
      <div class="flex-grow-1" />
      <v-menu
        offset-y
        v-if="userName"
      >
        <template v-slot:activator="{ on }">
          <v-btn
            text
            v-on="on"
          >
            {{ userName }}
          </v-btn>
        </template>
        <v-list>
          <v-list-item @click="logout">
            <v-list-item-title>ログアウト</v-list-item-title>
          </v-list-item>
        </v-list>
      </v-menu>
    </v-app-bar>

    <v-content>
      <router-view />
    </v-content>
  </v-app>
</template>

<script>
import { ROUTE_NAME } from '@/router/config';

export default {
  computed: {
    userName() {
      return this.$store.state.login.name;
    },
    routerToHomeName() {
      return ROUTE_NAME.HOME;
    },
  },
  methods: {
    logout() {
      this.$router.push('/logout');
    },
  },
};
</script>

<style scoped lang="scss">
.app {
  &__title {
    color: #fff;
    text-decoration: none;
  }
}
</style>
