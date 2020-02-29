<template>
  <div class="animated fadeIn">

    <b-row>
      <b-col sm="12">
        <c-table 
            :table-data="items" 
            :fields="fields" 
            :filterOn="filterOn"
            :per-page=10 
            hover striped bordered small fixed
            caption="<i class='fa fa-align-justify'></i> Switch List">
        </c-table>
      </b-col>
    </b-row>
   
  </div>

</template>

<script>
import cTable from '@/views/base/Table.vue'
import axios from 'axios'
export default {
  name: 'switchlist',
  components: {
      cTable
    },
  async mounted (){
      this.dataGet()
  },
  data: () => {
    return {
      items: [],
      apipath: '/api/v2/monitor/switch-controller/managed-switch/select',
      fields: [
            {key: 'name', label: 'Name', sortable: true},
            {key: 'serial', sortable: true},
            {key: 'state', sortable: true},
            {key: 'status', sortable: true},
            {key: 'os_version', sortable: true},
        ],
        filterOn: ['name','serial','state','status','os_version']
    }
  },
  methods: {
      dataGet: async function (){
            var response = await axios.get('http://localhost:4000/fosapi/FG140D/api/v2/monitor/switch-controller/managed-switch/select')
            this.items = response.data.results
        },
  }
}
</script>
