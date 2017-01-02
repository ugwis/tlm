<main>
	<form onsubmit={ export } >
		<button>export</button>
	</form>
	<adlibs adlibs={query.preparation.adlib} patch={adlibpatch} 
		showscreenname = { showscreenname }
		registerscreenname = { registerscreenname }
		/>
	<follower />
	<jobs jobs={ query.jobs } patch={ jobpatch }
		showscreenname = { showscreenname }
		registerscreenname = { registerscreenname }
	 />
	<submit />
	
	
	<script>
		query={
			preparation:{
				adlib:[],
				follower:[],
			},
			jobs:[

			],
			regularflag:false,
		};

		export(e) {
			console.log(JSON.stringify(query));
			return e.preventDefault();
		}

		adlibpatch(input) {
			query.preparation.adlib=input;
			this.update();
		}

		screenNameIdMap={};

		showscreenname(userID) {
			if (userID in screenNameIdMap ) {
				return screenNameIdMap[userID];
			} else {
				screenNameIdMap[userID]="";
				$.post("/api/getusers",{userids:userID},function(result){
					if(result.status==="ok"){
						screenNameIdMap[result.data[0][1]]=result.data[0][0];
					} else {
						delete screenNameIdMap[userID];
					}
				}).done(this.update).fail(()=>{
					delete screenNameIdMap[userID];
				})
			}
		}

		registerscreenname(userID, screenName){
			screenNameIdMap[userID]=screenName;
		}
		
		jobpatch(input) {
			query.jobs = input;
			this.update();
		}

	</script>
</main>

<adlibs>
	<h1>adlib</h1>
	<form onsubmit={ add } >
		tag:<input>
		<button>Add</button>
	</form>

	<li each={ v,i in opts.adlibs }>
		<adlib data={v} index={i} remove={ remove } patch={ patch }
		showscreenname={ parent.opts.showscreenname }
		registerscreenname={ parent.opts.registerscreenname }
		/>
	</li>

	<script>

		add(e) {
			var now = opts.adlibs;
			now.push({list:{tag:e.target[0].value},userids:[]});
			opts.patch(now);
			return e.preventDefault();
		}

		remove(index) {
			var now = opts.adlibs;
			now.splice(index, 1);
			opts.patch(now);
		}

		patch(index,input) {
			var now = opts.adlibs;
			now[index] = input;
			opts.patch(now);
		}
	</script>
</adlibs>

<adlib>
	<a onclick={ remove } class="glyphicon glyphicon-remove"></a>
	{opts.data.list.tag}<br />
	<span each={ i in opts.data.userids }>
		<a value={i} onclick={ delmember } class="glyphicon glyphicon-remove"></a>
		{parent.opts.showscreenname(i)}
	</span>
	<br />
	screenname:<input onBlur={fetch_user_lists} />
	<form style="display: inline" onsubmit={ addmember }>
		<select id="select" class="list-select" >
			<option each={selectlist} value={this[1]}>{this[0]}</option>
		</select>
		<button>Add</button>
	</form>


	<script>
		selectlist=[];

		remove() {
			opts.remove(opts.index);
		}

		addmember(e) {
			var now = opts.data;
			var user = e.target[0].value;
			if (user !== "" && now.userids.indexOf(user) === -1){
				now.userids.push(user);
				opts.patch(opts.index,now);
			}
			return 	e.preventDefault();
		}

		delmember(e) {
			var now = opts.data;
			var user = e.srcElement.value;
			var index = now.userids.indexOf(user);
			now.userids.splice(index, 1);
			opts.patch(opts.index,now);
		}

		fetch_user_lists(obj){
			screen_name = obj.srcElement.value;
			$.post("/api/searchuser",{username:screen_name},function(result){
				if(result.status==="ok"){
					this.selectlist = result.data;
					for(i in result.data) {
						opts.registerscreenname(result.data[i][1],result.data[i][0]);
					}
					this.update();
				}
			}.bind(this));
		}.bind(this);
	</script>
</adlib>

<follower>
	<h1>follower</h1>
</follower>

<jobs>
	<h1>Job</h1>
	<form onsubmit={ add } >
		<button>Add</button>
	</form>
	
	<li each={v,i in opts.jobs }>
		<job data={v} index={i} remove={ remove } patch={ patch }
		showscreenname={ parent.opts.showscreenname }
		registerscreenname={ parent.opts.registerscreenname }
		registerlistname={ registerlistname }
		/>
	</li>
	
	<script>
		listidname={};
	
		add(e){
			var now = opts.jobs;
			now.push({
				operator:"+",
				list1:{listid:0,tag:""},
				list2:{listid:0,tag:""},
				listresult:{listid:0,tag:""},
				config:{name:"",publicflag:false,saveflag:false},
				});
			opts.patch(now);
			return e.preventDefault();
		}
		
		patch(index, data){
			var now = opts.jobs;
			now[index] = data;
			opts.patch(now);
		}
		
		remove(index) {
			var now = opts.jobs;
			now.splice(index, 1);
			opts.patch(now);
		}
		
		registerlistname(listID, listName){
			listidname[listID]=listName;
		}
	</script>
</jobs>

<job>
	<a onclick={ remove } class="glyphicon glyphicon-remove"></a>
	<joblistselect data={ opts.data.list1 } patch={ list1patch } registerscreenname={opts.registerscreenname} registerlistname={opts.registerlistname} />
	<form onchange={ operatorchange }>
	<select id="operator" >
		<option value="+" selected={opts.data.operator==="+"}>+</option>
		<option value="*" selected={opts.data.operator==="*"}>*</option>
		<option value="-" selected={opts.data.operator==="-"}>-</option>
	</select>
	</form>
	<joblistselect data={ opts.data.list2 } patch={ list2patch } 
	registerscreenname={opts.registerscreenname} 
	registerlistname={opts.registerlistname} />
	=
	<joblistresult config={ opts.data.config } list={opts.data.listresult}
		configpatch={ configpatch } listpatch={ listpatch }
		registerscreenname={opts.registerscreenname} 
		registerlistname={opts.registerlistname}
	/>
	
	<script>
		remove() {
			opts.remove(opts.index);
		}
		
		list1patch(input) {
			var now = opts.data;
			now.list1=input;
			opts.patch(opts.index,now);
		}
		
		list2patch(input) {
			var now = opts.data;
			now.list2=input;
			opts.patch(opts.index,now);
		}
		
		operatorchange(e) {
			var now = opts.data;
			now.operator = e.target.value;
			opts.patch(now);
		}
		
		configpatch(input) {
			var now = opts.data;
			now.config=input;
			opts.patch(opts.index,now);
		}
		
		listpatch(input) {
			var now = opts.data;
			now.listresult=input;
			opts.patch(opts.index,now);
		}
		
	</script>
</job>

<joblistselect>
	<form onchange={ this.change }>
	<select id="select" >
		<option value="Tag" selected>Tag</option>
		<option value="ListID">List</option>
	</select>
	</form>
	
	<input if={this.type === "Tag"} onchange={ tagnamechange } name="test" value={opts.data.tag}>

	<selectlist if={this.type === "ListID"} changelistid={changelistid} 
		registerscreenname={opts.registerscreenname} 
		registerlistname={opts.registerlistname}
	/>
	
	<script>
		this.type = "Tag";
		selectlist=[];
		listlist=[];
	
		change(e) {
			this.type = e.target.value;
			this.update();
		}
		
		tagnamechange(e) {
			var now = opts.data;
			now.listid=0;
			now.tag=e.target.value;
			opts.patch(now);
		}
		
		changelistid(id) {
			var now = opts.data;
			now.tag="";
			now.listid = id;
			opts.patch(now);
		}
	</script>
</joblistselect>

<joblistresult>
	<form onchange={ this.change }>
	<select id="select" >
		<option value="NewSave" selected>NewSave</option>
		<option value="UpdateSave">UpdateSave</option>
		<option value="NotSave">NotSave</option>
	</select>
	</form>
	<div if={this.type==="NewSave"}>
		Name:<input onchange={namechange}>
		<form onchange={PrivatePubricChange}>
			<select id="PrivatePubric">
				<option value="Private">Private</option>
				<option value="Public">Public</option>	
			</select>
		</form>
	</div>
	
	<!-- loginしているユーザーリストに帰るべき -->
	<selectlist if={this.type==="UpdateSave"} changelistid={changelistid} 
		registerscreenname={opts.registerscreenname} 
		registerlistname={opts.registerlistname}
	/>
	
	<div if={this.type==="NotSave"}>
		Tag:<input  onchange={resulttagchange}>
	</div>
	
	<script>
		this.type="NewSave";
		change(e) {
			this.type = e.target.value;
			this.update();
		}
		
		PrivatePubricChange(e) {
			var confignow = opts.config;
			confignow.publicflag=(e.target.value==="Public");
			confignow.saveflag=true;
			opts.configpatch(confignow);
			var listnow = {
				ownerid:0,
				listid:0,
				tag:"",
			}
			opts.listpatch(listnow);
		}
		
		namechange(e) {
			var confignow = opts.config;
			confignow.name=e.target.value;
			confignow.saveflag=true;
			opts.configpatch(confignow);
			var listnow = {
				ownerid:0,
				listid:0,
				tag:"",
			}
			opts.listpatch(listnow);
		}
		
		changelistid(id) {
			var confignow = {
					name:"",
					saveflag:true,
					publicflag:false,
				}
			opts.configpatch(confignow);
			var listnow = {
				ownerid:0,
				listid:id,
				tag:"",
			}
			opts.listpatch(listnow);
		}
		
		resulttagchange(e) {
			var confignow = {
					name:"",
					saveflag:false,
					publicflag:false,
				}
			opts.configpatch(confignow);
			var listnow = {
				ownerid:0,
				listid:0,
				tag:e.target.value,
			}
			opts.listpatch(listnow);
		}
	</script>
</joblistresult>

<selectlist>
	screenname:<input id="screenname" onBlur={fetch_user_lists}>
	<select id="select" class="list-select" onBlur={fetch_lists}>
		<option each={selectlist} value={this[1]}>{this[0]}</option>
	</select>
	<select id="select" class="list-select" onBlur={changelistid}>
		<option each={listlist} value={this[1]}>{this[0]}</option>
	</select>
	
	<script>
		fetch_user_lists(obj){
			screen_name = obj.target.value;
			$.post("/api/searchuser",{username:screen_name},function(result){
				if(result.status==="ok"){
					this.selectlist = result.data;
					for(i in result.data) {
						opts.registerscreenname(result.data[i][1],result.data[i][0]);
					}
					this.update();
				}
			}.bind(this));
		}.bind(this);
		
		fetch_lists(obj) {
			userid = obj.target.value;
			$.post("/api/userlist",{userid:userid},function(result){
				if(result.status==="ok"){
					this.listlist = result.data;
					for(i in result.data) {
						opts.registerlistname(result.data[i][1],result.data[i][0]);
					}
					this.update();
				}
			}.bind(this));
		}.bind(this);
		
		changelistid(obj) {
			opts.changelistid(obj.target.value)
		}
	</script>
</selectlist>

<submit>
</submit>
