components {
  id: "pigeon"
  component: "/main/Scripts/pigeon.script"
  properties {
    id: "duration"
    value: "0.1"
    type: PROPERTY_TYPE_NUMBER
  }
  properties {
    id: "from_scale_animation"
    value: "0.12"
    type: PROPERTY_TYPE_NUMBER
  }
  properties {
    id: "to_scale_animation"
    value: "0.1"
    type: PROPERTY_TYPE_NUMBER
  }
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
