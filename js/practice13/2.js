function partial(firstArgument, secondArgument) {
	return function(thirdArgument, fourthArgument, fifthArgument) {
		return firstArgument + secondArgument + thirdArgument + fourthArgument + fifthArgument;
	}
}

function curry(firstArgument) {
	return function(secondArgument) {
		return function(thirdArgument) {
			return function(fourthArgument ) {
				return firstArgument + secondArgument + thirdArgument + fourthArgument;
			}
		}
	}
}