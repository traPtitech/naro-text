<template>
	<div>
		<div>ItemList</div>
		<ul>
			<li v-for="item in items" :key="item.name">
				<div>名前: {{ item.name }}</div>
				<div>{{ item.price }} 円</div>
			</li>
		</ul>
	</div>
</template>