components {
  id: "pigeon"
  component: "/main/Scripts/pigeon.script"
}
embedded_components {
  id: "sprite"
  type: "sprite"
  data: "default_animation: \"Pigeon\"\n"
  "material: \"/builtins/materials/sprite.material\"\n"
  "textures {\n"
  "  sampler: \"texture_sampler\"\n"
  "  texture: \"/main/Resources/Textures.atlas\"\n"
  "}\n"
  ""
  scale {
    x: 0.12
    y: 0.12
    z: 0.12
  }
}
