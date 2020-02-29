<template>
  <div class="animated fadeIn">

    <b-modal id="modalItem" :title="modalItem.title" @ok="modalOk" @hide="resetInfoModal">
      <b-row>
        <b-col sm="12">
          <b-form-group>
            <label for="name">Name</label>
            <b-form-input type="text" v-model="curitem.name" id="name"></b-form-input>
          </b-form-group>
        </b-col>
      </b-row>
      <b-row>
        <b-col sm="12">
          <b-form-group>
            <label for="device">Device</label>
            <b-form-select id="device" v-model="curitem.device" class="mb-3">
              <option v-for="dev in devicelist" v-bind:key="dev.name" :value="dev.name">{{dev.name}}</option>
            </b-form-select>
          </b-form-group>
        </b-col>
      </b-row>
      <b-row>
        <b-col sm="12">
          <b-form-group>
            <label for="endpoint">API Endpoint</label>
            <b-form-input type="text" v-model="curitem.endpoint" id="endpoint" ></b-form-input>
          </b-form-group>
        </b-col>
      </b-row>
      <b-row>
        <b-col sm="6">
          <b-form-group>
            <label for="methodget">Get Method</label>
            <b-form-input type="text" v-model="curitem.methodget" id="methodget" ></b-form-input>
          </b-form-group>
        </b-col>
        <b-col sm="6">
          <b-form-group>
            <label for="methodset">Set Method</label>
            <b-form-input type="text" v-model="curitem.methodset" id="methodset" ></b-form-input>
          </b-form-group>
        </b-col>
      </b-row>
      <b-row>
        <b-col sm="9"></b-col>
        <b-col sm="3">
          <b-form-group>
            <b-button variant="primary" type="text" @click="getSnippet" size="sm" id="methodget" >Get Snippet</b-button>
          </b-form-group>
        </b-col>
      </b-row>
      <b-row>
        <b-col sm="12">
          <b-form-group>
            <label for="devtype">Snippet</label>
            <b-form-textarea 
             v-model="curitem.snippet" 
             id="snippet"
             rows="3"
             max-rows="6"
              style="font-family:monospace;"
             ></b-form-textarea>
          </b-form-group>
        </b-col>
      </b-row>
     
    </b-modal>
    <b-row>
      <b-col sm="12">
        <crudtable 
            :table-data="items" 
            :fields="fields" 
            :filterOn="filterOn"
            :per-page=10 
            @row-edit="rowEdit"
            @row-delete="rowDelete"
            @row-dblclicked="rowEdit"
            @row-create="rowCreate"
            hover striped bordered small
            caption="<i class='fa fa-align-justify'></i> Snippet List"
          >

        </crudtable>
      </b-col>
    </b-row>
  </div>

</template>

<script>
import cTable from '@/views/base/Table.vue'
import crudtable from '@/views/base/CrudTable.vue'
import axios from 'axios'
export default {
  name: 'snippetlist',
  components: {
      cTable,
      crudtable
    },
  async mounted (){
       this.deviceGet()
  },
  data: () => {
    return {
      curitem: {},
      devicelist: [],
      items: [
        {name: 'snippet1', device: 'FGVMMAC', endpoint: '/api/v2/cmdb/system/interface', methodget: 'GET', methodset: 'POST',snippet: 'config system interface\nedit wan1\nappend allowaccess https ssh\nnext\nedit wan2\nappend allowaccess https ssh\nend\nend'}
      ],
      fields: [
            {key: 'name', label:'Name', sortable: true},
            {key: 'device',label:'Device', sortable: true},
            {key: 'endpoint',label:'API Endpoint', sortable: true},
            {key: 'actions', Label: 'Actions'}
        ],
        filterOn: ['name','device','endpoint'],
        modalItem: {
          id: 'info-modal',
          title: '',
          mode: '',
        }
    }
  },
  methods: {
      deviceGet: async function (){
            var response = await axios.get('http://localhost:4000/targets')
            this.devicelist = response.data
        },
      resetInfoModal() {
        this.modalItem.title = ''
        this.modalItem.content = ''
      },
      rowEdit(item){
        //console.log(item.name,item.firewallip,item.apikey)
        this.curitem = item
        this.modalItem.mode = 'edit'
        this.modalItem.title = 'Edit snippet'
        this.$root.$emit('bv::show::modal', "modalItem")
      },
      rowCreate(item){
        //console.log(item.name,item.firewallip,item.apikey)
        this.curitem = {}
        this.modalItem.mode = 'create'
        this.modalItem.title = 'Create snippet'
        this.$root.$emit('bv::show::modal', "modalItem")
      },
      modalOk(){
        if (this.modalItem.mode == 'edit'){
          this.apiUpdateSnippet()
        }
        if (this.modalItem.mode == 'create'){
          this.apiCreateSnippet()
        }
      },
      apiCreateSnippet: async function(){
        console.log(this.curitem)
        var response = await axios.post('http://localhost:4000/snippets',this.curitem)
        console.log(response)
        this.dataGet()
      },
      apiUpdateSnippet: async function(){
        console.log(this.curitem)
        var response = await axios.put('http://localhost:4000/snippets',this.curitem)
        console.log(response)
        this.dataGet()
      },
      rowDelete: async function(item){
        var response = await axios.delete('http://localhost:4000/snippets',{ data: item})
        console.log(item)
        this.dataGet()
      },
      getSnippet: async function(){
        console.log(this.curitem, this.curitem.endpoint, this.curitem.methodget, this.curitem.device)
        var response = await axios.post('http://localhost:4000/targets/'+this.curitem.device+'/api',{ 
                  path: this.curitem.endpoint,
                  method: this.curitem.methodget
                  })
        console.log(JSON.stringify(response.data.results))
        this.curitem.snippet = JSON.stringify(response.data.results, null, 2)
      },
  }
}
</script>
