module.exports = {
	newArray: function(initList, s, t)
	{
		return new GoArray(initList, s, t);
	},
}

//TODO: toString
function GoArray(initList, s, t)
{
	var arr = initList;
	var size = s;
	var type = t;

	this.getVal = function(index)
	{
		checkBounds(index);

		return arr;
	}

	this.getSlice = function(start, end)
	{
		if(end>=size || end==undefined)
		{
			end = size;
		}

		if(start<0 || start==undefined)
		{
			start=0;
		}

		return new Slice(start, end, this);
	}

	this.getStartPos = function()
	{
		return 0;
	}

	this.size = function()
	{
		return size;
	}


	function checkBounds(index)
	{
		if(index < 0 || index>=size)
		{
			throw "Index "+index+" is out of bounds."
		}
	}

	function getUndefined()
	{
		return 0;
	}
}


function Slice(start, end, origin)
{
	var capacity = origin.size() - start;
	var length = end-start;
	var startPos = start;

	this.getVal = function(index)
	{
		checkBounds(index)

		return origin.getVal(index+startPos);
	}

	this.set = function(index, val)
	{
		checkBounds(index)
		return origin.set(index+start, val);
	}

	this.getSlice = function(s, e)
	{
		return origin.getSlice(start+s, start+e);
	}

	this.getStartPos = function()
	{
		return startPos;
	}

	this.changeSlice = function(newStart, newEnd)
	{
		startPos = startPos+newStart;
		capacity = origin.size() - startPos;
		length = newEnd - startPos;
	}

	function checkBounds(index)
	{
		if(index<0 || index>=(end-start))
		{
			throw "Index "+index+" is out of bounds."
		}
	}
}

