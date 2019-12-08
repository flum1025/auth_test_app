<template>
  <div class="app-error">
    <div>
      <p>Error</p>
      <p>
        {{ code }} {{ message }}
      </p>
    </div>
  </div>
</template>

<script>
export default {
  props: {
    code: {
      type: Number,
      default: 500,
    },
    message: {
      type: String,
      default: 'Sorry, something went wrong',
    },
  },
};
</script>

<style scoped lang="scss">
.app-error {
  display: flex;
  height: 100%;
  width: 100%;
  justify-content: center;
  align-items: center;
  text-align: center;
  font-size: 2rem;

  p {
    margin-bottom: 0;
  }
}
</style>
