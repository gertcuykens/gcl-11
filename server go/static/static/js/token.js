Property = {
    key:null,
    value:null,
    add: function(t) {t.extra.push({key:this.key, value:this.value})}
}

Token = function() {
    this.id_token=null
    this.type_token=null
    this.access_token=null
    this.refresh_token=null
    this.expires_in=null
    this.expiry=new Date().toJSON()
    this.extra=[]
    this.status=null
}

Tokeng = function() {
    this.access_token=null
    this.error=null
    this.expires_in=null
    this.state=null
}
