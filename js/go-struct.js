module.exports = {
	defStruct : function(elementList){
		return new Struct(elementList)
	}
}

function Struct(elements){

	var elementList = elements

	this.getNew = function(parameters){
		var ret = {}

		if(Array.isArray(parameters)){
			//console.log("Array")
			elementList.forEach(function(e, index){
				ret[e] = parameters[index];
			})			
		}else{
			//console.log("Object")
			elementList.forEach(function(e, index){
				ret[e] = parameters[e]
			})	
		}

		return ret;
	}
}
