<template>
  <div class="animated fadeIn">
    <b-row>
      <b-col sm="12">
             <codemirror v-model="code" :options="cmOptions"></codemirror>
      </b-col>
    </b-row>
     <b-row>
      <b-col sm="12">
             <b-button @click="toconsole">clickme</b-button>
      </b-col>
    </b-row>
  </div>
</template>

<script>

import { codemirror } from 'vue-codemirror'
import 'codemirror/mode/go/go.js'
import 'codemirror/theme/base16-dark.css'
import axios from 'axios'
export default {
  name: 'floweditor',
  components: {
      codemirror
    },
  async mounted (){
      this.dataGet()
  },
  data: () => {
    return {
      editor: '',
      code: '',
      cmOptions: {
        tabSize: 4,
        mode: 'text/x-go',
        // theme: 'base16-dark',
        lineNumbers: true,
        line: true,
      }
    }
  },
  methods: {
      dataGet: async function (){
            var response = await axios.get('http://localhost:4000/targets')
            this.items = response.data
        },
      resetInfoModal() {
        this.modalItem.title = ''
        this.modalItem.content = ''
      },
      rowEdit(item){
        //console.log(item.name,item.firewallip,item.apikey)
        this.curitem = item
        this.modalItem.mode = 'edit'
        this.modalItem.title = 'Edit target'
        this.$root.$emit('bv::show::modal', "modalItem")
      },
      rowCreate(item){
        //console.log(item.name,item.firewallip,item.apikey)
        this.curitem = {}
        this.modalItem.mode = 'create'
        this.modalItem.title = 'Create target'
        this.$root.$emit('bv::show::modal', "modalItem")
      },
      modalOk(){
        if (this.modalItem.mode == 'edit'){
          this.apiCreateTarget()
        }
      },
      apiCreateTarget: async function(){
        console.log(this.curitem)
        var response = await axios.post('http://localhost:4000/targets',this.curitem)
        console.log(response)
        this.dataGet()
      },
      rowDelete: async function(item){
        var response = await axios.delete('http://localhost:4000/targets',{ data: item})
        console.log(item)
        this.dataGet()
      },
      toconsole(){
          console.log(this.code)
      }
  }
}
</script>
