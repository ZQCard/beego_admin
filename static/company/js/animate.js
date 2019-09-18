function fadeIn(node,callback){
	animate(node,'fadeIn',callback)
}
function animate(node,animationName,callback){
	node.classList.add('animated', animationName,'slow')

			function handleAnimationEnd() {
				node.classList.remove('animated', animationName,"action")
				node.removeEventListener('animationend', handleAnimationEnd)

				if (typeof callback === 'function') callback()
			}

		node.addEventListener('animationend', handleAnimationEnd)
}
