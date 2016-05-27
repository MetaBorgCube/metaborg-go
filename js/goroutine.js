function __go(func)
{
	func()
}

//TODO: Main Go Routine Handling
//TODO: Channels
//TODO: Deadlock Detection
//TODO: Stack Fixing
//TODO: Buffered Channels
//TODO: Non-Blocking Channels
//TODO: Think about panic
//TODO: Think about defer

var __scheduler = new Scheduler();

var __currentRoutine = new GoRoutine();

var mainroutine = 0;

function Scheduler()
{

	var scheduled = []
	var alive = false;

	this.schedule = function(routine, direct)
	{
		if(direct)
		{
			routine();
			return;
		}

		scheduled.push(routine);

		if(!alive)
		{
			this.start();
		}
	}

	this.start = function()
	{
		alive = true;

		//Use timeout, else this method would return then all GoRoutines are finished and the main 
		//would block  which would defeat the purpose of the 
		//Routine
		setTimeout(this.executeNext, 1);
	}

	this.executeNext = function()
	{
		try{
			var routine;

			while((routine = scheduled.shift()) != undefined)
			{
				routine();
			}

			alive = false;
		}finally{
			//If something goes wrong in a routine keep scheduling
			if(alive){
				setTimeout(this.executeNext, 0);
			}
		}

	}	
}

function GoRoutine()
{
	this.init = function(){

	}
}

function go(func)
{
	__scheduler.schedule(func);
}

