<template>
  <div>
    {{ query }}
  </div>
</template>

<script>
import QUERY from '@/graphql/query.gql';

export default {
  apollo: {
    query: {
      query: QUERY,
      update(data) {
        return data;
      },
      errorPolicy: 'all',
    },
  },
};
</script>
