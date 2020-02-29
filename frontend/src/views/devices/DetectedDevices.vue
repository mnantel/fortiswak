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
  name: 'DetectedDevices',
  components: {
      cTable
    },
  async mounted (){
      this.dataGet()
  },
  data: () => {
    return {
      items: [],
      fields: [
            {key: 'mac', sortable: true},
            {key: 'hostname', sortable: true},
            {key: 'ipv4_address', sortable: true},
            {key: 'is_online', sortable: true},
            {key: 'os_name', sortable: true},
            {key: 'sourcefg',sortable:true}
        ],
        filterOn: ['mac','hostname','ipv4_address','is_online','os_name','sourcefg']
    }
  },
  methods: {
      dataGet: async function (){
            var response = await axios.get('http://localhost:4000/monitor/user/detected-devices')
            this.items = response.data
        },
  }
}
</script>
