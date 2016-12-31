<main>
    <h3>Tag layout</h3>
	<adlib adlibs={query.adlib} patch={adlibpatch}/>
	<follower />
	<job />
	<submit />
	
	<script>
		query={
			adlib:[],
			};
		
		adlibpatch(input){
			query.adlib=input;
		}
	</script>
</main>

<adlib>
	<h1>Adlib</h1>
	<form onsubmit={ add }>
		<input>
		<button>Add</button>
	</form>
	
	<li each={ opts.adlibs }>
		{text}
	</li>
	
	<script>
	
		add(e){
			var now = opts.adlibs;
			now.push({text:e.target[0].value});
			opts.patch(now);
			return e.preventDefault(); 
		}
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