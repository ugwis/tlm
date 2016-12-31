<main>
	<form onsubmit={ export } >
		<button>export</button>
	</form>
	<adlibs adlibs={query.preparation.adlib} patch={adlibpatch} 
		showscreenname = {showscreenname }
		registerscreenname = { registerscreenname }
		/>
	<follower />
	<job />
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

	</script>
</main>

<adlibs>
	<h1>adlib</h1>
	<form onsubmit={ add } >
		tag:<input>
		<button>Add</button>
	</form>

	<li each={ opts.adlibs }>
		<adlib data={this} remove={ remove } patch={ patch }
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

		remove(event) {
			var now = opts.adlibs;
			var item = event.item;
			var index = now.indexOf(item);
			now.splice(index, 1);
			opts.patch(now);
		}

		patch(index,input) {
			var now = opts.adlibs;
			now[now.indexOf(index)] = input;
			opts.patch(now);
		}
	</script>
</adlibs>

<adlib>
	<a onclick={ opts.remove } class="glyphicon glyphicon-remove"></a>
	{opts.data.list.tag}<br />
	<span each={ i in opts.data.userids }>
		<a value={i} onclick={ delmember } class="glyphicon glyphicon-remove"></a>
		{parent.opts.showscreenname(i)}
	</span>
	<br />
	screenname:<input onkeydown={fetch_user_lists} />
	<form style="display: inline" onsubmit={ addmember }>
		<select id="select" class="list-select" >
			<option each={selectlist} value={this[1]}>{this[0]}</option>
		</select>
		<button>Add</button>
	</form>


	<script>
		selectlist=[];

		addmember(e) {
			var now = opts.data;
			var user = e.target[0].value;
			if (user !== "" && now.userids.indexOf(user) === -1){
				now.userids.push(user);
				opts.patch(opts.data,now);
			}
			return 	e.preventDefault();
		}

		delmember(e) {
			var now = opts.data;
			var user = e.srcElement.value;
			var index = now.userids.indexOf(user);
			now.userids.splice(index, 1);
			opts.patch(opts.data,now);
		}

		fetch_user_lists(obj){
			if(window.event.keyCode==13){
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
			}
		}.bind(this);
	</script>
</adlib>

<follower>
	<h1>follower</h1>
</follower>

<job>
	<h1>Job</h1>
</job>

<submit>
</submit>
