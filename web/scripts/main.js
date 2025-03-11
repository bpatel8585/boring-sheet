import * as THREE from 'three';
import WebGL from 'three/addons/capabilities/WebGL.js';
import { VRButton } from 'three/addons/webxr/VRButton.js';


const scene = new THREE.Scene();
const camera = new THREE.PerspectiveCamera( 75, window.innerWidth / window.innerHeight, 0.1, 1000 );

const renderer = new THREE.WebGLRenderer();
renderer.setSize( window.innerWidth, window.innerHeight );
document.body.appendChild( renderer.domElement );
document.body.appendChild( VRButton.createButton( renderer ) );

if ( WebGL.isWebGL2Available() ) {
	renderer.setAnimationLoop( function () {

		renderer.render( scene, camera );
	
	} );
} else {
	const warning = WebGL.getWebGL2ErrorMessage();
	document.getElementById( 'container' ).appendChild( warning );
}