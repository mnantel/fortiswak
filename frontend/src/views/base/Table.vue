<template>
  <b-card>
    <div slot="header" v-html="caption"></div>
    <b-row align-v="start">
      <b-col lg="5" class="my-1">
          <b-form-group
            label="Filter"
            label-cols-sm="2"
            label-align-sm="left"
            label-size="sm"
            label-for="filterInput"
            class="mb-0"
          >
            <b-input-group size="sm">
              <b-form-input
                v-model="filter"
                type="search"
                id="filterInput"
                placeholder="Type to Search"
              ></b-form-input>
              <b-input-group-append>
                <b-button :disabled="!filter" @click="filter = ''">Clear</b-button>
              </b-input-group-append>
            </b-input-group>
          </b-form-group>
      </b-col>
    </b-row>
    <b-table 
      responsive="sm" 
      :dark="dark" 
      :hover="hover" 
      :striped="striped" 
      :bordered="bordered" 
      :small="small" 
      :fixed="fixed" 
      :items="items" 
      :fields="captions" 
      :current-page="currentPage" 
      :per-page="perPage" 
      :filter="filter"
      :filterIncludedFields="filterOn"
      @row-clicked="rowClicked"
      @filtered="onFiltered"
      >
      <template slot="status" slot-scope="data">
        <b-badge :variant="getBadge(data.item.status)">{{data.item.status}}</b-badge>
      </template>
    </b-table>
    <nav>
      <b-pagination :total-rows="totalRows" :per-page="perPage" v-model="currentPage" prev-text="Prev" next-text="Next" hide-goto-end-buttons/>
    </nav>
  </b-card>
</template>

<script>


export default {
  name: 'c-table',
  inheritAttrs: false,
  props: {
    caption: {
      type: String,
      default: 'Table'
    },
    hover: {
      type: Boolean,
      default: false
    },
    striped: {
      type: Boolean,
      default: false
    },
    bordered: {
      type: Boolean,
      default: false
    },
    small: {
      type: Boolean,
      default: false
    },
    fixed: {
      type: Boolean,
      default: false
    },
    tableData: {
      type: [Array, Function],
      default: () => []
    },
    fields: {
      type: [Array, Object],
      default: () => []
    },
    perPage: {
      type: Number,
      default: 5
    },
    dark: {
      type: Boolean,
      default: false
    },
    filterOn: {
      type: [Array, String],
      default: () => []
    }
  },
  data: () => {
    return {
      currentPage: 1,
      filter:'',
      totalRows: 0
    }
  },
  // mounted() {
  //     this.totalRows = this.items.length
  // },
  watch:{
    items: function () {
      this.totalRows = this.items.length
    }
  },
  computed: {
    items: function() {
      const items =  this.tableData
      return Array.isArray(items) ? items : items()
    },
    // totalRows: function () { return this.items.length },
    captions: function() { return this.fields }
  },
  methods: {
    getBadge (status) {
      return status === 'Active' ? 'success'
        : status === 'Inactive' ? 'secondary'
          : status === 'Pending' ? 'warning'
            : status === 'Banned' ? 'danger' : 'primary'
    },
    getRowCount: function () {
      return this.items.length
    },
    rowClicked (item) {
      this.$emit('row-clicked', item)
    },
    onFiltered(filteredItems) {
        // Trigger pagination to update the number of buttons/pages due to filtering
        this.totalRows = filteredItems.length
        this.currentPage = 1
      }
  }
}
</script>
