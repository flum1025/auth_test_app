<template>
  <v-alert
    prominent
    type="error"
  >
    <v-row align="center">
      <v-col class="grow">
        {{ message }}
      </v-col>

      <slot />
    </v-row>
  </v-alert>
</template>

<script>
export default {
  props: {
    message: {
      type: String,
      required: true,
    },
  },
};
</script>
