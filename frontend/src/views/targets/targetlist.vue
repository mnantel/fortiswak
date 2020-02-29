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
            <label for="devtype">Device Type</label>
            <b-form-select id="devtype" v-model="curitem.devtype" class="mb-3">
              <option value="FortiGate">FortiGate</option>
              <option value="FortiManager">FortiManager</option>
            </b-form-select>
          </b-form-group>
        </b-col>
      </b-row>
      <b-row>
        <b-col sm="12">
          <b-form-group>
            <label for="firewallip">Device IP(:port)</label>
            <b-form-input type="text" v-model="curitem.firewallip" id="firewallip" ></b-form-input>
          </b-form-group>
        </b-col>
      </b-row>
      <b-row v-if="curitem.devtype === 'FortiManager'">
        <b-col sm="12">
          <b-form-group>
            <label for="username">Username</label>
            <b-form-input type="text" v-model="curitem.username" id="username" ></b-form-input>
          </b-form-group>
        </b-col>
      </b-row>
      <b-row>
        <b-col sm="12">
          <b-form-group>
            <label v-if="curitem.devtype === 'FortiManager'" for="apikey">Password</label>
            <label v-else for="apikey">API Key</label>
            <b-form-input type="password" v-model="curitem.apikey" id="apikey" ></b-form-input>
          </b-form-group>
        </b-col>
      </b-row>
      
      <b-row>
        <b-col sm="12">
          <b-form-group>
            <b-form-radio-group id="basicInlineRadios"
                  :plain="true"
                  v-model="curitem.enabled"
                  :options="[
                    {text: 'Enabled',value: 'true'},
                    {text: 'Disabled',value: 'false'},
                  ]"
                  >
            </b-form-radio-group>
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
            caption="<i class='fa fa-align-justify'></i> Target List"
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
  name: 'targetlist',
  components: {
      cTable,
      crudtable
    },
  async mounted (){
      this.dataGet()
  },
  data: () => {
    return {
      curitem: {},
      items: [],
      fields: [
            {key: 'name', label:'Name', sortable: true},
            {key: 'firewallip',label:'Management IP:Port', sortable: true},
            {key: 'devtype',label:'Device Type', sortable: true},
            {key: 'enabled',label:'Enabled', sortable: true},
            {key: 'status',label:'Status', sortable: true},
            {key: 'actions', Label: 'Actions'}
        ],
        filterOn: ['name','firewallip','enabled','status','devtype'],
        modalItem: {
          id: 'info-modal',
          title: '',
          mode: '',
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
          this.apiUpdateTarget()
        }
        if (this.modalItem.mode == 'create'){
          this.apiCreateTarget()
        }
      },
      apiCreateTarget: async function(){
        console.log(this.curitem)
        var response = await axios.post('http://localhost:4000/targets',this.curitem)
        console.log(response)
        this.dataGet()
      },
      apiUpdateTarget: async function(){
        console.log(this.curitem)
        var response = await axios.put('http://localhost:4000/targets',this.curitem)
        console.log(response)
        this.dataGet()
      },
      rowDelete: async function(item){
        var response = await axios.delete('http://localhost:4000/targets',{ data: item})
        console.log(item)
        this.dataGet()
      }
  }
}
</script>
