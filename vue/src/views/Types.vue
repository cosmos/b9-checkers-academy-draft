<template>
	<div>
		<div class="container">
			<h2>
			Create Game Module
			</h2>
			<CreateGame />

			<!-- this line is used by starport scaffolding # 4 -->
		</div>
	</div>
</template>

<script>
import CreateGame from "../components/CreateGame";
export default {
	name: 'Types',
	components: { CreateGame },
}
</script>
