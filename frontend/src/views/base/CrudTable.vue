<template>
  <b-card>
    <div slot="header" v-html="caption"></div>
    <b-row align-v="start" align-h="between">
      <b-col lg="6" class="my-1">
        <b-button size="sm" @click="rowCreate">Create</b-button>
      </b-col>
      <b-col lg="4" class="my-1">
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
      head-variant="dark"
      :filterIncludedFields="filterOn"
      @row-clicked="rowClicked"
      @filtered="onFiltered"
      selectable
      >
      <template v-slot:cell(enabled)="data">
        <b-badge :variant="getBadge(data.item.enabled)">{{data.item.enabled}}</b-badge>
      </template>
      <template v-slot:cell(actions)="row">
          <b-button size="sm" @click="rowEdit(row.item)" class="mr-1">
            Edit
          </b-button>
          <b-button size="sm" variant="danger" @click="rowDelete(row.item)" class="mr-1">
            Delete
          </b-button>
          <!-- <b-button size="sm" @click="row.toggleDetails">
            {{ row.detailsShowing ? 'Hide' : 'Show' }} Details
          </b-button> -->
      </template>
       
    </b-table>
    <nav>
      <b-pagination :total-rows="totalRows" :per-page="perPage" v-model="currentPage" prev-text="Prev" next-text="Next" hide-goto-end-buttons/>
    </nav>
  </b-card>
</template>

<script>


export default {
  name: 'CrudTable',
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
      return status === 'true' ? 'success'
        : status === 'false' ? 'warning' : 'primary'
    },
    getRowCount: function () {
      return this.items.length
    },
    rowClicked (item) {
      this.$emit('row-clicked', item)
    },
    rowEdit (item) {
      this.$emit('row-edit', item)
    },
    rowDelete (item) {
      this.$emit('row-delete', item)
    },
    rowCreate (item) {
      this.$emit('row-create', item)
    },
    onFiltered(filteredItems) {
        // Trigger pagination to update the number of buttons/pages due to filtering
        this.totalRows = filteredItems.length
        this.currentPage = 1
      }
  }
}
</script>
