class FirewallFilterFormView extends KDFormViewWithFields

  constructor: (options = {}, data) ->

    options.fields     =
      type             :
        label          : "Type"
        name           : "type"
        cssClass       : "half"
        itemClass      : KDSelectBox
        selectOptions  : [
          { title      : "IP"                 ,  value : "ip"             }
          { title      : "Country"            ,  value : "country"        }
          { title      : "Request per second" ,  value : "request.second" }
          { title      : "Request per minute" ,  value : "request.minute" }
        ]
        validate       :
          rules        :
            required   : yes
          messages     :
            required   : "Please select a filter type"
      value            :
        label          : "Value"
        name           : "value"
        cssClass       : "half"
        validate       :
          rules        :
            required   : yes
          messages     :
            required   : "Please select a filter type"
      remove           :
        itemClass      : KDCustomHTMLView
        cssClass       : "delete-button half"
        click          : =>
          @emit "FirewallFilterRemoved"
          @destroy()

    unless options.removable
      delete options.fields.remove
      options.cssClass = KD.utils.curry "full", options.cssClass

    super options, data
