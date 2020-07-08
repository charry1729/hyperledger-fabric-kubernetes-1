<template lang="pug">

  v-row
    v-col(md="8" offset-md="2")

        v-data-table(
        :header="headers"
        :items="raw_resources"
      )
        template(v-slot:item.weight="{item}")
          span {{item.weight}} kgs.

        template(v-slot:item.id="{item}")
          router-link(:to="'/update/'+ item.id") Update
          a(@click="destroy(item.id)") Destroy
          //- span {{item.id}}
</template>

<script>
// @ is an alias to /src
//import HelloWorld from '@/components/HelloWorld.vue'
import {mapState} from 'vuex';

export default {
  data(){
    return{
      
      headers:[
        {  text: "Name", value:'name'   },
        {  text: "Weight", value:'weight'   },
        {  text: "TimeStamp", value:'timestamp'   },
        {  text: "Actions", value:'id'   },
      ],


    };
     
  },
  computed:{
    ...mapState({
      raw_resources: state => state.raw_resources
    }),
  },
  methods: {

    async destroy(id){
      const vm =  this;
      await vm.$store.dispatch('destroy',id);
    }
  }

// name: 'Home',
//   components: {
//     HelloWorld
//   }
}
</script>
