<template>
	<div>
		<div>ItemList</div>
		<ul>
			<li
				v-for="item in items"
				:key="item.name"
				:class="{ over500: item.price >= 500 }"
			>
				<div>名前: {{ item.name }}</div>
				<div>{{ item.price }} 円</div>
				<div v-if="item.price >= 10000">高額商品</div>
			</li>
		</ul>
		==略==
	</div>
</template>