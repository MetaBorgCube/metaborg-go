module.exports = {
	newPointer: function(set, get)
	{
		return new Pointer(set, get);
	},
}

function Pointer(_set, _get)
{

	var obj;

	this.set = _set
	this.get = _get
}

