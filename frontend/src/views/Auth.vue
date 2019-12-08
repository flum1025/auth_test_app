<template>
  <AlertBox
    class="auth__error"
    v-if="hasError"
    message="Authentication failure"
  >
    <v-col class="shrink">
      <v-btn
        :to="{ name: routerToHomeName }"
      >
        LOGIN
      </v-btn>
    </v-col>
  </AlertBox>
  <div v-else>
    authorizing...
  </div>
</template>

<script>
import { handleAuthentication } from '@/services/auth';
import { ACTION_TYPES } from '@/store/modules/login';
import { ROUTE_NAME } from '@/router/config';

import AlertBox from '@/components/AlertBox.vue';

export default {
  components: { AlertBox },
  data() {
    return {
      hasError: false,
    };
  },
  computed: {
    routerToHomeName() {
      return ROUTE_NAME.HOME;
    },
  },
  async created() {
    try {
      const authResult = await handleAuthentication();
      this.$store.dispatch(`login/${ACTION_TYPES.logIn}`, authResult);
      this.$router.push({ name: ROUTE_NAME.HOME });
    } catch (e) {
      console.error(e);
      this.hasError = true;
    }
  },
  render() {
    return null;
  },
};
</script>

<style lang="scss">
.auth {
  &__error {
    margin: 10px;
  }
}
</style>
