export default {
  items: [
    // {
    //   name: 'Dashboard',
    //   url: '/dashboard',
    //   icon: 'icon-speedometer',
    //   badge: {
    //     variant: 'primary',
    //     text: 'NEW'
    //   }
    // },
    {
      name: 'Snippets',
      url: '/snippets',
      icon: 'fa fa-cog',
    },
    {
      name: 'Admin',
      url: '/admin',
      icon: 'fa fa-cog',
      children: [
        {
          name: 'Target List',
          url: '/admin/targetlist',
          icon: 'fa fa-list'
        },
     
      ]
    },
  ]
}
