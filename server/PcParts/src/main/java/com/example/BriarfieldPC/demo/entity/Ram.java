package com.example.BriarfieldPC.demo.entity;

import lombok.AllArgsConstructor;
import lombok.Builder;
import lombok.Data;
import lombok.NoArgsConstructor;

import javax.persistence.Entity;
import javax.persistence.GeneratedValue;
import javax.persistence.Id;

import static javax.persistence.GenerationType.IDENTITY;

@Data
@AllArgsConstructor
@NoArgsConstructor
@Builder
@Entity
public class Ram {
    @Id
    @GeneratedValue(strategy = IDENTITY)
    private long id;
    private Integer memory;
    private String clock;
    private Integer TDP;
    private Integer price;
}
