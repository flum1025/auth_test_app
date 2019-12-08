<template>
  <div>
    bye bye
  </div>
</template>

<script>
import { ACTION_TYPES } from '@/store/modules/login';

export default {
  mounted() {
    this.$store.dispatch(`login/${ACTION_TYPES.logOut}`);
  },
};
</script>
