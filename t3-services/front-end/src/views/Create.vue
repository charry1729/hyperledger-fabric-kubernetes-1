<template lang="pug">
    v-row 
        v-col(md="6" offset-md="2")
            v-text-field(
                label="Name"
                v-model="name"
            )

            v-text-field(
                label="Weight"
                v-model="weight"
            )
            v-select(
                label="Type"
                v-model="type_id"
                :items="raw_resource_types"
                item-text="name"
                item-value="id"
            )

            v-btn(color="success" @click="create") Create

</template>

<script>
import {mapState} from 'vuex';
//import { mapState } from 'vuex'
export default {
    name: Create,
    data(){
        return{
            name: "",
            weight: 42000,
            type_id : 1,

        }
    },
    computed : {
        ...mapState({
        raw_resource_types: state => state.raw_resources_types
        })
        
    },
    methods: {
        async create() {
            const vm = this;
            await vm.$store.dispatch('create',{
                name: vm.name,
                weight: vm.weight,
                type_id: vm.type_id
            });
        }
    }
}
</script>