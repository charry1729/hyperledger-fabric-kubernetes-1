<template lang="pug">
  div
    v-toolbar
      v-toolbar-title Welcome

      v-toolbar-items
        v-btn(text to="/") Home
        v-btn(text to="/create") Create
    v-app
      v-container(fluid)
      router-view
</template>

<script>
//import HelloWorld from './components/HelloWorld';

export default {
  name: 'App',

  // components: {
  //   HelloWorld,
  // },

  data: () => ({
    //
  }),
};
</script>

<style scoped>
  .wrapper{
    padding: 20px
  }
 
</style>
