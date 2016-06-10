module.exports = {
	go: go,
	newChannel: function(size)
	{
		return new Channel(size);
	},
	end: function()
	{
		__scheduler.unregister();
	}
}

//TODO: Main Go Routine Handling 
//TODO: Registration
//TODO: Deadlock Detection
//TODO: Stack Fixing
//TODO: Work stealing?
//TODO: Think about panic
//TODO: Think about defer

var __scheduler = new Scheduler();
var nrSleepingRoutines = 0;


var mainroutine = 0;

function Scheduler()
{

	//TODO: channellist for Deadlock detection
	var scheduled = []
	var alive = false;

	var totalRoutines = 1;

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

		//Use timeout, else this method would return when all GoRoutines are finished and the main 
		//would block  which would defeat the purpose of the 
		//Routine
		setTimeout(this.executeNext, 1);
	}

	this.register = function()
	{
		totalRoutines++;
	}

	this.unregister = function()
	{
		totalRoutines--;
	}

	this.executeNext = function()
	{
		try{
			var routine;

			while((routine = scheduled.shift()) != undefined)
			{
				routine();
			}

			if(nrSleepingRoutines == totalRoutines)
			{
				console.log("All Go Routines are sleeping, deadlock detected.")
			}

			//TODO: Insert Deadlock Detection here
			alive = false;
		}finally{
			//If something goes wrong in a routine keep scheduling
			if(alive){
				setTimeout(this.executeNext, 0);
			}
		}

	}	
}

function Channel(size)
{
	var buffer = [];
	var buffersize = size 

	//TODO: Check multiple waiting
	var waitingReceivers = [];
	var waitingSenders = [];

	this.send = function(message, callback)
	{
		if(buffer.length <= buffersize)
		{
			if(waitingReceivers.length > 0)
			{
				var recCallback = waitingReceivers.shift()

				//Release waiting routine 
				scheduleReceiver(recCallback, message)
				nrSleepingRoutines--;
			}else{
				buffer.push(message);
			}
			
			__scheduler.schedule(callback)
		}else{
			nrSleepingRoutines++;
			waitingSenders.push({
				data: message,
				callback: callback
			});
		}
	}

	this.receive = function(callback)
	{
		var data = buffer.shift();
		var sender = null;

		if(waitingSenders.length > 0)
		{
			sender = waitingSenders.shift();
			__scheduler.schedule(sender.callback)
			nrSleepingRoutines--;
		}

		if(data !== undefined)
		{
			scheduleReceiver(callback, data)

			//If the buffer was full, 
			//senders could be waiting to send
			if(sender){
				 buffer.push(sender.data)
			}
		}else{
			if(sender)
			{
				scheduleReceiver(callback, sender.data)
			}else{

				//Wait for message
				waitingReceivers.push(callback);
				nrSleepingRoutines++;
			}
		}
	}

	function scheduleReceiver(callback, data)
	{
		__scheduler.schedule(function(){
			callback(data);
		})
	}

	this.getNrOfWaitingRoutines = function()
	{
		return waitingSenders.length + waitingReceivers.length;
	}


} 

function go(func)
{
	__scheduler.register();
	__scheduler.schedule(func);
}

