# Example Dashboard JSON
resource "datadog_dashboard_json" "dashboard_json" {
  dashboard = <<EOF
{
   "title":"Ordered Layout Dashboard",
   "description":"Created using the Datadog provider in Terraform",
   "widgets":[
      {
         "id":719369537777170,
         "definition":{
            "title":"Widget Title",
            "type":"alert_graph",
            "alert_id":"895605",
            "viz_type":"timeseries"
         }
      },
      {
         "id":2306240030393868,
         "definition":{
            "title":"Widget Title",
            "type":"alert_value",
            "alert_id":"895605",
            "unit":"b",
            "text_align":"center",
            "precision":3
         }
      },
      {
         "id":6990998850881326,
         "definition":{
            "title":"Widget Title",
            "type":"alert_value",
            "alert_id":"895605",
            "unit":"b",
            "text_align":"center",
            "precision":3
         }
      },
      {
         "id":3351284044659007,
         "definition":{
            "title":"Widget Title",
            "type":"change",
            "requests":[
               {
                  "q":"avg:system.load.1{env:staging} by {account}",
                  "compare_to":"week_before",
                  "change_type":"absolute",
                  "order_dir":"desc",
                  "increase_good":true,
                  "order_by":"name",
                  "show_present":true
               }
            ]
         }
      },
      {
         "id":6450290622996182,
         "definition":{
            "title":"Widget Title",
            "show_legend":false,
            "type":"distribution",
            "requests":[
               {
                  "q":"avg:system.load.1{env:staging} by {account}",
                  "style":{
                     "palette":"warm"
                  }
               }
            ]
         }
      },
      {
         "id":4902842646291536,
         "definition":{
            "title":"Widget Title",
            "type":"check_status",
            "check":"aws.ecs.agent_connected",
            "grouping":"cluster",
            "group_by":[
               "account",
               "cluster"
            ],
            "tags":[
               "account:demo",
               "cluster:awseb-ruthebdog-env-8-dn3m6u3gvk"
            ]
         }
      },
      {
         "id":6392349954822644,
         "definition":{
            "title":"Widget Title",
            "show_legend":false,
            "type":"heatmap",
            "yaxis":{
               "scale":"sqrt",
               "include_zero":true,
               "min":"1",
               "max":"2"
            },
            "requests":[
               {
                  "q":"avg:system.load.1{env:staging} by {account}",
                  "style":{
                     "palette":"warm"
                  }
               }
            ]
         }
      },
      {
         "id":5222961478940988,
         "definition":{
            "title":"Widget Title",
            "type":"hostmap",
            "requests":{
               "fill":{
                  "q":"avg:system.load.1{*} by {host}"
               },
               "size":{
                  "q":"avg:memcache.uptime{*} by {host}"
               }
            },
            "node_type":"container",
            "no_metric_hosts":true,
            "no_group_hosts":true,
            "group":[
               "host",
               "region"
            ],
            "scope":[
               "region:us-east-1",
               "aws_account:727006795293"
            ],
            "style":{
               "palette":"yellow_to_green",
               "palette_flip":true,
               "fill_min":"10",
               "fill_max":"20"
            }
         }
      },
      {
         "id":8121199734227072,
         "definition":{
            "type":"note",
            "content":"note text",
            "background_color":"pink",
            "font_size":"14",
            "text_align":"center",
            "show_tick":true,
            "tick_pos":"50%",
            "tick_edge":"left"
         }
      },
      {
         "id":1775856835833038,
         "definition":{
            "title":"Widget Title",
            "type":"query_value",
            "requests":[
               {
                  "q":"avg:system.load.1{env:staging} by {account}",
                  "aggregator":"sum",
                  "conditional_formats":[
                     {
                        "hide_value":false,
                        "comparator":"<",
                        "palette":"white_on_green",
                        "value":2
                     },
                     {
                        "hide_value":false,
                        "comparator":">",
                        "palette":"white_on_red",
                        "value":2.2
                     }
                  ]
               }
            ],
            "autoscale":true,
            "custom_unit":"xx",
            "text_align":"right",
            "precision":4
         }
      },
      {
         "id":8461455966625581,
         "definition":{
            "title":"Widget Title",
            "type":"query_table",
            "requests":[
               {
                  "q":"avg:system.load.1{env:staging} by {account}",
                  "aggregator":"sum",
                  "conditional_formats":[
                     {
                        "hide_value":false,
                        "comparator":"<",
                        "palette":"white_on_green",
                        "value":2
                     },
                     {
                        "hide_value":false,
                        "comparator":">",
                        "palette":"white_on_red",
                        "value":2.2
                     }
                  ],
                  "limit":10
               }
            ]
         }
      },
      {
         "id":8660006349418736,
         "definition":{
            "title":"Widget Title",
            "type":"scatterplot",
            "requests":{
               "x":{
                  "q":"avg:system.cpu.user{*} by {service, account}",
                  "aggregator":"max"
               },
               "y":{
                  "q":"avg:system.mem.used{*} by {service, account}",
                  "aggregator":"min"
               }
            },
            "xaxis":{
               "scale":"pow",
               "label":"x",
               "include_zero":true,
               "min":"1",
               "max":"2000"
            },
            "yaxis":{
               "scale":"log",
               "label":"y",
               "include_zero":false,
               "min":"5",
               "max":"2222"
            },
            "color_by_groups":[
               "account",
               "apm-role-group"
            ]
         }
      },
      {
         "id":1669590772917638,
         "definition":{
            "title":"env: prod, datacenter:us1.prod.dog, service: master-db",
            "title_size":"16",
            "title_align":"left",
            "type":"servicemap",
            "service":"master-db",
            "filters":[
               "env:prod",
               "datacenter:us1.prod.dog"
            ]
         }
      },
      {
         "id":2138829058361817,
         "definition":{
            "title":"Widget Title",
            "show_legend":true,
            "legend_size":"2",
            "type":"timeseries",
            "requests":[
               {
                  "q":"avg:system.cpu.user{app:general} by {env}",
                  "on_right_yaxis":false,
                  "metadata":[
                     {
                        "expression":"avg:system.cpu.user{app:general} by {env}",
                        "alias_name":"Alpha"
                     }
                  ],
                  "style":{
                     "palette":"warm",
                     "line_type":"dashed",
                     "line_width":"thin"
                  },
                  "display_type":"line"
               },
               {
                  "on_right_yaxis":false,
                  "log_query":{
                     "index":"mcnulty",
                     "search":{
                        "query":"status:info"
                     },
                     "group_by":[
                        {
                           "facet":"host",
                           "sort":{
                              "facet":"@duration",
                              "aggregation":"avg",
                              "order":"desc"
                           },
                           "limit":10
                        }
                     ],
                     "compute":{
                        "facet":"@duration",
                        "interval":5000,
                        "aggregation":"avg"
                     }
                  },
                  "display_type":"area"
               },
               {
                  "on_right_yaxis":false,
                  "apm_query":{
                     "index":"apm-search",
                     "search":{
                        "query":"type:web"
                     },
                     "group_by":[
                        {
                           "facet":"resource_name",
                           "sort":{
                              "facet":"@string_query.interval",
                              "aggregation":"avg",
                              "order":"desc"
                           },
                           "limit":50
                        }
                     ],
                     "compute":{
                        "facet":"@duration",
                        "interval":5000,
                        "aggregation":"avg"
                     }
                  },
                  "display_type":"bars"
               },
               {
                  "on_right_yaxis":false,
                  "process_query":{
                     "search_by":"error",
                     "metric":"process.stat.cpu.total_pct",
                     "limit":50,
                     "filter_by":[
                        "active"
                     ]
                  },
                  "display_type":"area"
               }
            ],
            "yaxis":{
               "scale":"log",
               "include_zero":false,
               "max":"100"
            },
            "events":[
               {
                  "q":"sources:test tags:1"
               },
               {
                  "q":"sources:test tags:2"
               }
            ],
            "markers":[
               {
                  "label":" z=6 ",
                  "value":"y = 4",
                  "display_type":"error dashed"
               },
               {
                  "label":" x=8 ",
                  "value":"10 < y < 999",
                  "display_type":"ok solid"
               }
            ]
         }
      },
      {
         "id":7307171374656551,
         "definition":{
            "title":"Widget Title",
            "type":"toplist",
            "requests":[
               {
                  "q":"avg:system.cpu.user{app:general} by {env}",
                  "conditional_formats":[
                     {
                        "hide_value":false,
                        "comparator":"<",
                        "palette":"white_on_green",
                        "value":2
                     },
                     {
                        "hide_value":false,
                        "comparator":">",
                        "palette":"white_on_red",
                        "value":2.2
                     }
                  ]
               }
            ]
         }
      },
      {
         "id":7086674838553258,
         "definition":{
            "title":"Group Widget",
            "type":"group",
            "layout_type":"ordered",
            "widgets":[
               {
                  "id":3726092277657502,
                  "definition":{
                     "type":"note",
                     "content":"cluster note widget",
                     "background_color":"pink",
                     "font_size":"14",
                     "text_align":"center",
                     "show_tick":true,
                     "tick_pos":"50%",
                     "tick_edge":"left"
                  }
               },
               {
                  "id":6376384650558057,
                  "definition":{
                     "title":"Alert Graph",
                     "type":"alert_graph",
                     "alert_id":"123",
                     "viz_type":"toplist"
                  }
               }
            ]
         }
      },
      {
         "id":4668903563678912,
         "definition":{
            "title":"Widget Title",
            "type":"slo",
            "view_type":"detail",
            "time_windows":[
               "7d",
               "previous_week"
            ],
            "slo_id":"56789",
            "show_error_budget":true,
            "view_mode":"overall",
            "global_time_target":"0"
         }
      }
   ],
   "template_variables":[
      {
         "name":"var_1",
         "default":"aws",
         "prefix":"host"
      },
      {
         "name":"var_2",
         "default":"autoscaling",
         "prefix":"service_name"
      }
   ],
   "layout_type":"ordered",
   "is_read_only":true,
   "notify_list":[

   ],
   "template_variable_presets":[
      {
         "name":"preset_1",
         "template_variables":[
            {
               "name":"var_1",
               "value":"host.dc"
            },
            {
               "name":"var_2",
               "value":"my_service"
            }
         ]
      }
   ]
}
EOF
}
