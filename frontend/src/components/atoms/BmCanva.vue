<template>


  <v-stage :config="stageSize" class="w-full" @wheel="handleWheel" ref="stageRef">
    <v-layer>
      <v-image
        v-if="imageShow"
        :config="{
          image: imageShow.value,
          width: imageSize.width,
          height: imageSize.height,
          x: imageSize.x,
          y: imageSize.y,
        }"
      />
    </v-layer>
    <v-layer  v-if="polys">
       <v-group v-for="poly in polys" >
        <v-line   :config="polygonConfig(poly)"/>
        <v-circle @dragmove="handleDragMove" @dragend="handleDragEnd" v-for="point in polygonPoints(poly)" :config="pointConfig(point)" />
      </v-group>
    </v-layer>
  </v-stage>
</template>

<script setup>
import { watch, ref } from 'vue';
import { useImage } from 'vue-konva';
import { useTokenStore } from '../../store.js'
import { useActionState } from 'react';
import BmIconButton from './BmIconButton.vue';

const tokenStore = useTokenStore()

const props = defineProps({
  imageURL: {
    type: String,
    required: true,
    immediate: true
  },
  polygons: {
    type: Array,
    default: () => [],
    immediate: true
  },
  projectID: {
    type: String
  }
});

const stageRef = ref(null);

const handleWheel = (e) => {
  e.evt.preventDefault();

  const stage = stageRef.value.getNode();
  const oldScale = stage.scaleX();
  const pointer = stage.getPointerPosition();

  const mousePointTo = {
    x: (pointer.x - stage.x()) / oldScale,
    y: (pointer.y - stage.y()) / oldScale,
  };

  // how to scale? Zoom in? Or zoom out?
  let direction = e.evt.deltaY > 0 ? 1 : -1;

  // when we zoom on trackpad, e.evt.ctrlKey is true
  // in that case lets revert direction
  if (e.evt.ctrlKey) {
    direction = -direction;
  }

  const scaleBy = 1.1;
  const newScale = direction > 0 ? oldScale * scaleBy : oldScale / scaleBy;

  stage.scale({ x: newScale, y: newScale });

  const newPos = {
    x: pointer.x - mousePointTo.x * newScale,
    y: pointer.y - mousePointTo.y * newScale,
  };
  stage.position(newPos);
};

let classes = {}

  fetch(`http://localhost:8000/api/project/${props.projectID}/classes`, {
    method: "GET", // Or POST, PUT, DELETE, etc.
    headers: {
        "Authorization": `Bearer ${tokenStore.token}`,
        "Content-Type": "application/json" // Example for JSON data
    },
  })
  .then(response => response.json()) // Parse the response as JSON      
  .then(data => {
    for (let classe of data.data) {
      classes[classe.id] = classe
    }
  })
  .catch(error => {
    console.error('Error fetching data:', error);
  })


const polys = ref(parsePolygons(props.polygons));

const stageSize = ref({
  width: window.innerWidth,
  height: window.innerHeight,
  draggable: true
});

const imageSize = ref({
  width: window.innerWidth,
  height: window.innerHeight,
  x: 0,
  y: 0
})

const [image] = useImage(props.imageURL);
const imageShow = ref(image);

function getImageDimensions(imageUrl) {
  return new Promise((resolve, reject) => {
    const img = new Image();

    img.onload = () => {
      resolve({
        width: img.naturalWidth,
        height: img.naturalHeight
      });
    };

    img.onerror = () => {
      reject(new Error('Failed to load image or image not found.'));
    };

    img.src = imageUrl;
  });
}

watch(() => props.imageURL, (newURL) => {
 
  getImageDimensions(newURL)
    .then(dimensions => {
      imageSize.value.x = 0
      imageSize.value.y = 0
      imageSize.value.width = dimensions.width;
      imageSize.value.height = dimensions.height;

      const [img] = useImage(newURL);
      imageShow.value = img
    })
    .catch(error => {
      console.error(error.message);
    })
  
});

watch(() => props.polygons, (newPolygons) => {
  polys.value = parsePolygons(props.polygons )
});

function parsePolygons(polygons) {
  console.log(polygons)
  
  let output = []   

  for (let poly of polygons) {
    let coords = []

    for (let ponto of poly.coordinates) {
      coords.push(ponto[0])
      coords.push(ponto[1])
    }

    output.push({
      id: poly.id,
      classe: classes[poly.classID] , 
      coords,
    })

  }
  
  return output
}

function polygonPoints(poly) {
  let output = []
  let counter = 0

  for (let i = 1; i < poly.coords.length; i+=2) {
    output.push({ id: `${poly.id}@${counter}`, color: poly.classe?.color, x: poly.coords[i - 1], y: poly.coords[i] })
    counter++
  }
  return output
}

function polygonConfig(poly) {
  console.log(poly)
  return {
    id: poly.id,
    points: poly.coords,
    fill: poly.classe?.color + '88',
    stroke: poly.classe?.color,
    strokeWidth: 1,
    closed: true,
  };
}

function pointConfig(point) {
  return {
    id: point.id,
    x: point.x,
    y: point.y,
    fill: point.color,
    radius: 4,
    stroke: 'black',
    strokeWidth: 1,
    closed: true,
    draggable: true
  };
}        

const handleDragMove = (event) => {
  const [polyID, pointID] = event.target.attrs.id.split('@')
  polys.value.map(poly => {
    if (poly.id == polyID) {
      poly.coords[2*pointID] = event.target.attrs.x
      poly.coords[2*pointID + 1] = event.target.attrs.y
    }
    return poly
  })
}

const handleDragEnd = (event) => {
  const [polyID, pointID] = event.target.attrs.id.split('@')

    const poly =  props.polygons.find(poly => {
    if (poly.id == polyID) {
      return poly
    }
  })

  const polyINFO = polys.value.find(poly => {
    if (poly.id == polyID) {
      return poly
    }
  })

  console.log(polyINFO)
  
  const coords = []

  for (let i = 1; i < polyINFO.coords.length; i+=2) {
    coords.push([ polyINFO.coords[i - 1], polyINFO.coords[i]])
  }

  fetch(`http://localhost:8000/api/project/image/${poly.imageID}/polygon/${poly.id}`, {
    method: "PUT", // Or POST, PUT, DELETE, etc.
    headers: {
        "Authorization": `Bearer ${tokenStore.token}`,
        "Content-Type": "application/json" // Example for JSON data
    },
    body: JSON.stringify({
        name: poly.name,
        classID: poly.classID,
        coordinates: coords,
    })
  })
  .then(response => response.json()) // Parse the response as JSON      
  .then(data => {
    console.log(data)
  })
  .catch(error => {
    console.error('Error fetching data:', error);
  })
}



</script>
